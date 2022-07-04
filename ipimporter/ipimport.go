package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"io/ioutil"
	"net/http"
	"net/netip"
	"os"
)

type GoogleCloudFile struct {
	SyncToken    string `json:"syncToken"`
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}

type AmazonWebServicesFile struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix string `json:"ip_prefix"`
	} `json:"prefixes"`
}

type Ipfile struct {
	Url              string
	DownloadFilePath string
	CloudPlatform    string
}

func (i *Ipfile) Download() (err error) {
	//Download the IP Address file
	txt := fmt.Sprintf("Downloading %s", i.Url)
	fmt.Println(txt)

	// Create the file
	fileOut, err := os.Create(i.DownloadFilePath)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	resp, err := http.Get(i.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Errorf("bad status: %s", resp.Status))
	}

	// Write the body to file
	_, err = io.Copy(fileOut, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func googleAsJson(DownloadFilePath string) (fileOut GoogleCloudFile) {
	// Open downloaded file and return as json
	jsonFile, _ := os.Open(DownloadFilePath)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &fileOut)

	return fileOut
}

func amazonAsJson(DownloadFilePath string) (fileOut AmazonWebServicesFile) {
	// Open downloaded file and return as json

	jsonFile, _ := os.Open(DownloadFilePath)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &fileOut)

	return fileOut
}
func ExpandCidr(cidr string) (ips []netip.Addr, err error) {
	//parse a cidr and return all the ip addresses
	prefix, err := netip.ParsePrefix(cidr)
	ip_addr := prefix.Addr()
	for {
		ips = append(ips, ip_addr)
		ip_addr = ip_addr.Next()
		if prefix.Contains(ip_addr) == false {
			break
		}
	}
	return ips, err
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Str_in_slice(str string, slice []string) bool {
	//find a string in slice return boolean
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
