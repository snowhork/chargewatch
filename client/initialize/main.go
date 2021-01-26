package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var (
		host        = flag.String("host", "http://localhost:8088", "Server host")
		user        = flag.String("user", "", "Your user name (this is unique)")
		name        = flag.String("name", "", "Your device name")
		description = flag.String("description", "", "Your device description")
	)
	flag.Parse()

	registerDevice(*host, *user, *name, *description)
}

type RequestPayload struct {
	Name        string
	Description string
}

func registerDevice(host, user, name, description string) {
	payload := RequestPayload{Name: name, Description: description}
	raw, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/user/%s/devices", host, user),
		bytes.NewBuffer(raw),
	)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	println(string(res))
}
