package main

import (
	"errors"
	"net"
)

func ipv4toDecimal(ipIn net.IP) (decimalOut int, err error) {
	//Convert an IP4 Address to a decimal
	ipOct := net.IP.To4(ipIn)

	if ipOct == nil {
		err := errors.New("Error Coverting to IP")
		return 0, err
	}

	octInts := [4]int{int(ipOct[0]) * 16777216, int(ipOct[1]) * 65536, int(ipOct[2]) * 256, int(ipOct[3])}

	for _, value := range octInts {
		decimalOut = decimalOut + value
	}
	return decimalOut, nil
}
