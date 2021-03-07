package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// TracerouteData stores info needed to perform a Traceroute
type TracerouteData struct {
	Host string `json:"host"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// Traceroute does a traceroute
func Traceroute(w http.ResponseWriter, r *http.Request) {
	// Get data from original request to form the traceroute command
	body, _ := ioutil.ReadAll(r.Body)
	data := TracerouteData{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Failed to unmarshal json into params")
	}
}
