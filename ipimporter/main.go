package main

import (
	"context"
	"fmt"
	"log"
	"net/netip"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/sync/semaphore"
)

func ipfiles() []Ipfile {

	googleCloudPlatform := Ipfile{
		DownloadFilePath: "google-cloud.json",
		Url:              "https://www.gstatic.com/ipranges/cloud.json",
		CloudPlatform:    "Google",
	}

	amazonWebServices := Ipfile{
		DownloadFilePath: "aws-ips.json",
		Url:              "https://ip-ranges.amazonaws.com/ip-ranges.json",
		CloudPlatform:    "aws",
	}

	googleGeneral := Ipfile{
		DownloadFilePath: "goog.json",
		Url:              "https://www.gstatic.com/ipranges/goog.json",
		CloudPlatform:    "google",
	}

	return []Ipfile{googleGeneral, amazonWebServices, googleCloudPlatform}
}

func main() {

	MONGOHOST := os.Getenv("MONGOHOST")
	if MONGOHOST == "" {
		MONGOHOST = "localhost"
	}

	var (
		client   *mongo.Client
		mongoURL = fmt.Sprintf("mongodb://mongo:mongo@%s:27017", MONGOHOST)
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	collection_ips := client.Database("ips").Collection("col_ips")
	for _, ipfile := range ipfiles() {
		ipfile.Download()

		var jsonObj interface{}
		var cidrs []string

		if ipfile.CloudPlatform == "aws" {
			jsonObj = amazonAsJson(ipfile.DownloadFilePath)
			json := jsonObj.(AmazonWebServicesFile)
			fmt.Printf("Found %v Cidrs from %s", len(json.Prefixes), ipfile.CloudPlatform)
			for _, val := range json.Prefixes {
				exists := Str_in_slice(val.IPPrefix, cidrs)
				if exists == false {
					cidrs = append(cidrs, val.IPPrefix)
				}
			}
		}

		if ipfile.CloudPlatform == "google" {
			jsonObj = googleAsJson(ipfile.DownloadFilePath)
			json := jsonObj.(GoogleCloudFile)
			fmt.Printf("Found %v Cidrs from %s", len(json.Prefixes), ipfile.CloudPlatform)
			for _, val := range json.Prefixes {
				var cidr string
				if len(val.Ipv4Prefix) > 0 {
					cidr = val.Ipv4Prefix
				}
				// } else {
				// 	cidr = val.Ipv6Prefix
				// }
				exists := Str_in_slice(cidr, cidrs)

				if exists == false {
					cidrs = append(cidrs, cidr)
				}
			}
		}

		//Create a slice to expand all the IP addresses that are valid to each Cidr.
		// 81.10.20.1 , 81.10.20.2, 81.10.20.3 etc

		var ipSubnets [][]netip.Addr

		for _, cidr := range cidrs {
			ipSubnet, err := ExpandCidr(cidr)
			if err != nil {
				log.Println(err)
			}
			ipSubnets = append(ipSubnets, ipSubnet)
		}

		sem := semaphore.NewWeighted(15)

		//Add the IP addresses into mongo
		for _, sn := range ipSubnets {
			sem.Acquire(context.Background(), 1)
			subnet := CreateSubnetBatch(sn, ipfile.Url, ipfile.CloudPlatform)
			fmt.Printf("Inserting %v IP addresses from %s\n", len(subnet), ipfile.CloudPlatform)
			go func() {
				InsertMany(subnet, collection_ips)
				sem.Release(1)
			}()
		}
	}

}
