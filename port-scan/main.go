package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello, I am a custom port scanner!")
	fmt.Println("Please enter the host and port you want to scan.....")
	// fmt.Println("Args: ", os.Args[1], os.Args[2])
	hostName := flag.String("host", "example.com", "The host you want to scan")
	portNumber := flag.String("port", "", "The port you want to scan")

	flag.Parse()
	host := *hostName
	port := *portNumber

	// fmt.Println("Host: ", host)
	// fmt.Println("Port: ", port)

	if port != "" {
		conn, err := net.Dial("tcp", host+":"+"1433")
		if err != nil {
			fmt.Println("Error while connecting to host: ", err)
			return
		} else {
			fmt.Println("Port open: ", 1433)
		}
		conn.Close()
	} else {
		fmt.Println("Scanning host: ", host)
		for port := 1; port < 1024; port++ {
			// fmt.Println("Scanning port: ", port)
			portScan(host, port)
		}
	}
}

func portScan(host string, port int) {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		// fmt.Println("Port closed: ", port)
		if oppErr, ok := err.(*net.OpError); ok {
			if oppErr.Timeout() {
				fmt.Printf("Timeout error: on %v . %v", port, err)
			} else {
				// Port is closed
				// fmt.Printf("Error while connecting to port %v . %v: ", port, err)
				// fmt.Printf("Port closed: %v\n", port)
			}
		}
	} else {
		fmt.Println("Port open: ", port)
		defer conn.Close()
	}
}
