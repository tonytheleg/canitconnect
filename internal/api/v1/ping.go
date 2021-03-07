package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// PingData stores info needed to perform a Ping
type PingData struct {
	Host string `json:"host"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// Ping pings a host
func Ping(w http.ResponseWriter, r *http.Request) {
	// Get data from original request to form the traceroute command
	body, _ := ioutil.ReadAll(r.Body)
	data := PingData{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Failed to unmarshal json into params")
	}
}
