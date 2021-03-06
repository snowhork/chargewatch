package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type RequestPayload struct {
	ChargeValue int  `json:"chargeValue"`
	Charging    bool `json:"charging"`
}

func SendCharge(host, deviceID string) {
	r, err := exec.Command("system_profiler", "SPPowerDataType").Output()
	if err != nil {
		log.Fatal(err)
	}

	var rawRemain, rawFull, rawCharging string
	for _, line := range strings.Split(string(r), "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Charge Remaining") {
			rawRemain = strings.Split(line, " ")[3]
			continue
		}
		if strings.Contains(line, "Full Charge Capacity") {
			rawFull = strings.Split(line, " ")[4]
			continue
		}
		if strings.Contains(line, "Charging") {
			rawCharging = strings.Split(line, " ")[1]
			continue
		}
	}

	remain, err := strconv.Atoi(string(rawRemain))
	if err != nil {
		log.Fatal(err)
	}
	full, err := strconv.Atoi(string(rawFull))
	if err != nil {
		log.Fatal(err)
	}

	value := 100*remain/(full+1) + 1
	charging := string(rawCharging) == "Yes"

	if err := post(host, deviceID, value, charging); err != nil {
		log.Println(err)
	}
}

func post(host, deviceID string, value int, charging bool) error {
	payload := RequestPayload{ChargeValue: value, Charging: charging}
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/devices/%s/charge", host, deviceID)

	println(string(path))
	println(string(raw))

	req, err := http.NewRequest(
		"POST",
		path,
		bytes.NewBuffer(raw),
	)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	println(string(res))

	return nil
}
