package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"sync"
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

	var wg sync.WaitGroup

	for i := 0; i < len(ports); i++ {
		wg.Add(1)
		go scann(address, ports[i], &wg)
	}

	wg.Wait()
}

func scann(address string, port int, wg *sync.WaitGroup){
	addressPort := address + ":" + strconv.Itoa(port)
	connection, err := net.DialTimeout("tcp", addressPort, 5*time.Second)
	if err == nil {
		fmt.Printf("Port %d in %s is OPEN UwU\n", port, address)
		connection.Close()
	} else {
		fmt.Printf("Port %d in %s is CLOSED or timeout UnU\n", port, address)
	}
	wg.Done()
}
