package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	address := "scanme.nmap.org"
	port := 80
	addressPort := address + ":" + strconv.Itoa(port)
	_, err := net.Dial("tcp", addressPort)
	if err == nil {
		fmt.Printf("Port %d in %s is open UwU\n", port, address)
	} else {
		fmt.Printf("Port %d in %s is close UnU", port, address)
	}
}
