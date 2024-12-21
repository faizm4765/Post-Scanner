package main

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

func main() {
	println("Hello, I am a custom port scanner!")
	pinger, err := probing.NewPinger("example.com")
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()
	fmt.Printf("Ping packets sent: %d\n", stats.PacketsSent)
	fmt.Printf("Ping packets received: %d\n", stats.PacketsRecv)
	fmt.Printf("Ping packets lost: %d\n", stats.PacketsSent-stats.PacketsRecv)
}
