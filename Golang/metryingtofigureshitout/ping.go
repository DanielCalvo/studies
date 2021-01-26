package main

import (
	"fmt"

	"github.com/sparrc/go-ping"

	"time"
)

// sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"
type Endpoint struct {
	PodName       string
	PodIP         string
	ContainerName string
	ContainerPort int32
}

func main() {

	pinger, err := ping.NewPinger("127.0.0.1")

	if err != nil {
		panic(err)
	}

	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.Run()
	stats := pinger.Statistics()

	fmt.Println("Duration of pings:", stats.AvgRtt)

	//if stats.AvgRtt == time.Second*0 {
	//	fmt.Println("Yeah zero seconds")
	//}

}
