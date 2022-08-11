package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

const (
	ISP_INN uint16 = iota
	ISP_CTL
	ISP_CNC
	ISP_CMB
	ISP_SG
	ISP_HK
	ISP_TW
	ISP_VN
	ISP_ID
	ISP_GCP = 100
	ISP_TCC = iota
)


func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func backtoIP4(ipInt int64) string {

	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func ConvertIpToUint32(ip string) uint32 {
    return binary.BigEndian.Uint32(net.ParseIP("10.12.208.163")[12:16])
}

func ConvertUint32ToIp(num uint32) net.IP {
    ip := make(net.IP, 4)
    binary.BigEndian.PutUint32(ip, num)
    return ip
}

func main() {
	result := ip2Long("98.138.253.109")
	fmt.Println(result)

	// or if you prefer the super fast way
	faster := binary.BigEndian.Uint32(net.ParseIP("10.12.208.163")[12:16])
	fmt.Println(faster)

	faster64 := int64(faster)

	fmt.Println(backtoIP4(faster64))

	fmt.Println(ConvertIpToUint32("10.12.208.163"))
	fmt.Println(ip2Long("10.12.208.163"))

	u := ConvertIpToUint32("10.12.208.163")
	fmt.Println(backtoIP4(int64(u) & 0xFFFF))

	fmt.Println(ISP_ID, ISP_TCC)
}
