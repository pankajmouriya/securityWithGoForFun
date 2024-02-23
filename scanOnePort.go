package main

import (
	"fmt"
	"net"
)

func scanOnePort() {
	con, err := net.Dial("tcp", "noshellaccess.com:80")
	if err == nil {
		fmt.Println("Connection Successful")
		fmt.Println(con.RemoteAddr())
	}
}
