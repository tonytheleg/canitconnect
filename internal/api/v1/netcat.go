package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"text/template"
	"time"
)

var netcattpl *template.Template

func init() {
	netcattpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

// NetcatInput stores info needed to perform a netcat
type NetcatInput struct {
	Hostname string `json:"host"`
	Port     string `json:"port"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

type NetcatOutput struct {
	Hostname string
	Response string
}

// Netcat takes the passed URL and Port to test and mimics a netcat call to test if a port is open
func Netcat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data := NetcatInput{}
	if r.URL.String() == "/api/v1/netcat" {
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal json into params")
		}
		NetcatAPI(w, r, data)
	} else {
		data.Hostname = r.FormValue("hostname")
		data.Port = r.FormValue("port")
		NetcatForm(w, r, data)
	}
}

func NetcatAPI(w http.ResponseWriter, r *http.Request, data NetcatInput) {
	address := net.JoinHostPort(data.Hostname, data.Port)
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

func NetcatForm(w http.ResponseWriter, r *http.Request, data NetcatInput) {
	out := &NetcatOutput{}
	out.Hostname = data.Hostname

	address := net.JoinHostPort(data.Hostname, data.Port)
	timeout := time.Second * 10
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		out.Response = fmt.Sprint("Connection failed:", err)
	}
	if conn != nil {
		defer conn.Close()
		out.Response = fmt.Sprint("Connected to ", address, " (success!)")
	}
	result := Results{NetcatResp: out}
	err = netcattpl.ExecuteTemplate(w, "result.html", result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
