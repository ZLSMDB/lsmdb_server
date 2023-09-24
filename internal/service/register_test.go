package service

import (
	"fmt"
	"net"
	"testing"
)

func TestGetIP(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range addrs {
		// 检查是否是IP地址，并且不是回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("IPv4 Address:", ipnet.IP.String())
			}
		}
	}
	fmt.Println(fmt.Scanf("%s:%v"))
}
