package v1

import (
	"bytes"
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

	var b bytes.Buffer

	header := fmt.Sprintf("traceroute to %v %v hops max, %v byte packets\n", data.Host, options.MaxHops(), options.PacketSize())
	b.WriteString(header)
	c := make(chan traceroute.TracerouteHop, 0)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(&b, hop)
		}
	}()

	_, err = traceroute.Traceroute(data.Host, &options, c)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Fprintln(w, b.String())

}

func printHop(b *bytes.Buffer, hop traceroute.TracerouteHop) {
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	if hop.Success {
		hopString := fmt.Sprintf("%-3d %v (%v)  %v\n", hop.TTL, hostOrAddr, addr, hop.ElapsedTime)
		//fmt.Printf("%-3d %v (%v)  %v\n", hop.TTL, hostOrAddr, addr, hop.ElapsedTime)
		b.WriteString(hopString)
	} else {
		hopString := fmt.Sprintf("%-3d *\n", hop.TTL)
		//fmt.Printf("%-3d *\n", hop.TTL)
		b.WriteString(hopString)
	}
}

func address(address [4]byte) string {
	return fmt.Sprintf("%v.%v.%v.%v", address[0], address[1], address[2], address[3])
}
