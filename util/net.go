package util

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// 读数据
func ReadLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return l, nil
}

// 发送数据
func SendData(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}
	_, e := conn.Write(append([]byte(fmt.Sprintf("%d ", len(value))), value...))
	return e
}

func GetIPAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get IPv4 address fail", err)
		return "", err
	}
	var ipAddr *net.IPNet
	for _, addr := range addrs {
		// 检查是否是IP地址，并且不是回环地址
		if ipAddr, ok := addr.(*net.IPNet); ok && !ipAddr.IP.IsLoopback() {
			if ipAddr.IP.To4() != nil {
				fmt.Println("IPv4 Address:", ipAddr.IP.String())
			}
		}
	}
	return ipAddr.IP.String(), nil
}
