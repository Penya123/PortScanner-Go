package main

import (
	"fmt"
	"net"
)

func main() {
	address := ""
	port := ""
	addressPort := address + ":" + port
	_, err := net.Dial("tcp", addressPort)
	if err == nil {
		fmt.Println("Connection Successful UwU")
	} else {
		fmt.Println("Connection failed UnU")
	}
}
