package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	address := "scanme.nmap.org"
	var ports []int = []int{
		21,    //ftp
		22,    //ssh
		23,    //telnet
		25,    //SMTP
		53,    //DNS
		80,    //HTTP
		110,   //POP3
		143,   //IMAP
		443,   //HTTPS
		445,   //SMB
		993,   //IMAPS
		995,   //POP32
		3306,  //MySQL
		3389,  //RDP
		5432,  //PostgreSQL
		5900,  //VNC
		6379,  //Redis
		8080,  //proxy/dev
		8443,  //HTTPS alt
		27017, //MongoDB
	}

	for i := 0; i < len(ports); i++ {
		addressPort := address + ":" + strconv.Itoa(ports[i])
		connection, err := net.DialTimeout("tcp", addressPort, 15*time.Second)
		if err == nil {
			fmt.Printf("Port %d in %s is OPEN UwU\n", ports[i], address)
			connection.Close()
		} else {
			fmt.Printf("Port %d in %s is CLOSED or timeout UnU\n", ports[i], address)
		}
	}

}
