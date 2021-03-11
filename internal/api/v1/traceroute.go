package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aeden/traceroute"
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

	options := traceroute.TracerouteOptions{}
	options.SetRetries(1)
	options.SetMaxHops(64)
	options.SetFirstHop(1)

	fmt.Fprintf(w, "traceroute to %v %v hops max, %v byte packets\n", data.Host, options.MaxHops(), options.PacketSize())
	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(hop)
		}
	}()

	_, err = traceroute.Traceroute(data.Host, &options, c)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func printHop(hop traceroute.TracerouteHop) {
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	if hop.Success {
		fmt.Printf("%-3d %v (%v)  %v\n", hop.TTL, hostOrAddr, addr, hop.ElapsedTime)
	} else {
		fmt.Printf("%-3d *\n", hop.TTL)
	}
}

func address(address [4]byte) string {
	return fmt.Sprintf("%v.%v.%v.%v", address[0], address[1], address[2], address[3])
}
