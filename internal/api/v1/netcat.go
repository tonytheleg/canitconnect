package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// NetcatData stores info needed to perform a netcat
type NetcatData struct {
	Host string `json:"host"`
	Port string `json:"port"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// Netcat takes the passed URL and Port to test and mimics a netcat call to test if a port is open
func Netcat(w http.ResponseWriter, r *http.Request) {
	// Get data from original request to form the curl command
	body, _ := ioutil.ReadAll(r.Body)
	data := NetcatData{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Failed to unmarshal json into params")
	}
	address := net.JoinHostPort(data.Host, data.Port)
	timeout := time.Second * 10
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Fprintln(w, "Connection Failed:", err)
	}
	if conn != nil {
		defer conn.Close()
		fmt.Fprintln(w, "Connected to", address, "(success)!")
	}
}
