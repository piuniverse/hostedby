package main

import (
	"context"
	"fmt"
	"log"
	"net/netip"

	"go.mongodb.org/mongo-driver/mongo"
)

type bsonIp struct {
	Id            string `bson:"_id" json:"_id,omitempty"`
	Ip            string `bson:"ip_address" json:"ip_address"`
	Url           string `bson:"url" json:"url"`
	CloudPlatform string `bson:"cloudplatform" json:"cloudplatform"`
}

func InsertOne(ip bsonIp, mongoCollection *mongo.Collection) {
	//Insert One Record
	result, err := mongoCollection.InsertOne(context.TODO(), ip)
	// check for errors in the insertion
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(result.InsertedID)
	}
}

func InsertMany(ips []interface{}, mongoCollection *mongo.Collection) {
	//Insert Multiple Records
	_, err := mongoCollection.InsertMany(context.TODO(), ips)
	// check for errors in the insertion
	if err != nil {
		log.Println(err)
	} else {
		//
	}
}

func CreateSubnetBatch(IpsIn []netip.Addr, Url string, cloudPlatform string) (IpsOut []interface{}) {
	//Create a batch of subnets to post ready to Insert multiple records
	for _, val := range IpsIn {
		i := bsonIp{
			Id:            val.String(),
			Ip:            val.String(),
			Url:           Url,
			CloudPlatform: cloudPlatform,
		}
		IpsOut = append(IpsOut, i)
	}

	return IpsOut
}
