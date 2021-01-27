package main

import (
	"chargewatch/client/lib"
	"flag"
	"time"
)

func main() {
	var (
		host     = flag.String("host", "http://localhost:8088", "Server host")
		deviceID = flag.String("deviceID", "", "Device ID")
		interval = flag.Int("interval", 600, "interval (second)")
	)
	flag.Parse()

	ticker := time.NewTicker(time.Second * time.Duration(*interval))
	lib.SendCharge(*host, *deviceID)

	for {
		<-ticker.C
		lib.SendCharge(*host, *deviceID)
	}
}
