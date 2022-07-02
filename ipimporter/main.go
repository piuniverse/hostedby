package main

import (
	"fmt"
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

	return []Ipfile{googleGeneral, googleCloudPlatform, amazonWebServices}
}

func main() {

	for _, ipfile := range ipfiles() {
		ipfile.Download()

		var jsonObj interface{}
		var ips []IpObj

		if ipfile.CloudPlatform == "aws" {
			jsonObj = amazonAsJson(ipfile.DownloadFilePath)
			json := jsonObj.(AmazonWebServicesFile)
			for _, val := range json.Prefixes {
				ips = append(ips, IpObj{
					Ip:            val.IPPrefix,
					Url:           ipfile.Url,
					Type:          "IPv4",
					CloudPlatform: ipfile.CloudPlatform,
				})
			}
		}

		if ipfile.CloudPlatform == "google" {
			jsonObj = googleAsJson(ipfile.DownloadFilePath)
			json := jsonObj.(GoogleCloudFile)
			for _, val := range json.Prefixes {
				var IpAddr string
				var IpType string
				if len(val.Ipv4Prefix) > 0 {
					IpAddr = val.Ipv4Prefix
					IpType = "IPv4"
				} else {
					IpAddr = val.Ipv6Prefix
					IpType = "IPv6"
				}
				ips = append(ips, IpObj{
					Ip:            IpAddr,
					Url:           ipfile.Url,
					Type:          IpType,
					CloudPlatform: ipfile.CloudPlatform,
				})
			}
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

}
