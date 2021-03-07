package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CurlData stores info needed to perform a curl
type CurlData struct {
	URL string `json:"url"`
	//	HTTPProxy  string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func Curl(w http.ResponseWriter, r *http.Request) {
	// Get data from original request to form the curl command
	body, _ := ioutil.ReadAll(r.Body)
	data := CurlData{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Failed to unmarshal json into params")
	}
	// Create the new request
	req, err := http.NewRequest("GET", data.URL, nil)
	if err != nil {
		// handle err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(w, "Something failed", err)
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "%v %v\ncontent-length: %v\nserver: %v\ncontent-type: %v\ndata: %v\n",
		resp.Proto,
		resp.Status,
		resp.ContentLength,
		resp.Header["Server"][0],
		resp.Header["Content-Type"][0],
		resp.Header["Date"][0],
	)
}
