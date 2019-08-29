package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"time"
)

// sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"

func main() {
	pinger, err := ping.NewPinger("192.168.1.11")
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)

	fmt.Println("Duration of pings:", stats.AvgRtt)

	if stats.AvgRtt == time.Second*0 {
		fmt.Println("Yeah zero seconds")
	}

}
