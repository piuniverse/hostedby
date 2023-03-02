package main

import (
	"log"
	"net"
	"testing"
)

func TestIPv4toDecimal(t *testing.T) {
	//Test IPv4 to Decimal func
	ip, _, err := net.ParseCIDR("82.12.162.1/32")
	if err != nil {
		log.Panicln("ParseCIDR Error:", err)
	}
	got, _ := ipv4toDecimal(ip)
	want := 1376559617

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
