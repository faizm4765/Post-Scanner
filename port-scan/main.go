package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
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
		for port := 1; port < 65536; port++ {
			port := strconv.Itoa(port)
			// fmt.Println("Scanning port: ", port)
			conn, err := net.Dial("tcp", host+":"+string(port))
			if err != nil {
				// fmt.Println("Port closed: ", port)
				continue
			}

			fmt.Println("Port open: ", port)
			conn.Close()
		}
	}
}
