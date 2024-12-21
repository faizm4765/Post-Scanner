package main

import "github.com/go-ping/ping"

func main() {
	println("Hello, I am a custom port scanner!")
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()
	println(stats)
}
