package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello, I am a custom port scanner!")
	fmt.Println("Please enter the host and port you want to scan.....")
	// fmt.Println("Args: ", os.Args[1], os.Args[2])
	hostName := flag.String("host", "example.com", "The host you want to scan")
	portNumber := flag.String("port", "443", "The port you want to scan")

	flag.Parse()
	host := *hostName
	port := *portNumber

	fmt.Println("Host: ", host)
	fmt.Println("Host: ", port)
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error while connecting to host: ", err)
		return
	}

	fmt.Println("Scanning host: ", host)
	fmt.Println("Port open: ", port)
	conn.Close()
}
