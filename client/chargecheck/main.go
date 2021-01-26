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
	done := make(chan struct{})
	go func() {
		lib.SendCharge(*host, *deviceID)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				lib.SendCharge(*host, *deviceID)
			}
		}
	}()

	time.Sleep(100 * time.Second)
	ticker.Stop()
	done <- struct{}{}
}
