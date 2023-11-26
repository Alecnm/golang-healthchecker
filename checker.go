package main

import (
	"fmt"
	"net"
	"time"
)

func healthCheck(destination string, port string) string {
	address := destination + ":" + port
	timeOut := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeOut)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error : %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}