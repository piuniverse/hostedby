package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/sync/semaphore"
	"gopkg.in/yaml.v2"
)

func ipfiles() []Ipfile {
	//Load the IP files
	var ipfiles []Ipfile

	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("failed reading config file: %v\n", err)
	}

	err = yaml.Unmarshal(source, &ipfiles)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return ipfiles
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

		sem := semaphore.NewWeighted(1)
		for _, cidr := range cidrs {
			sem.Acquire(context.Background(), 1)
			expandedcidr, err := ExpandCidr(cidr)
			if err != nil {
				log.Println(err)
			}
			subnet := CreateSubnetBatch(expandedcidr, ipfile.Url, ipfile.CloudPlatform)
			expandedcidr = nil
			fmt.Printf("Inserting %v IP addresses from %s\n", len(subnet), ipfile.Url)
			go func() {
				InsertMany(subnet, collection_ips)
				sem.Release(1)
				subnet = nil
			}()
		}
		ipfile = Ipfile{}
	}
}
