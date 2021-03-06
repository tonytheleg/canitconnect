package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

var traceTpl *template.Template

func init() {
	traceTpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

// Traceroute takes the passed hostname and returns a traceroute to it
func Traceroute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	data := TracerouteInput{}
	if r.URL.String() == "/api/v1/traceroute" {
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal json into params")
		}
		TracerouteAPI(w, r, data)
	} else {
		data.Hostname = r.FormValue("hostname")
		TracerouteForm(w, r, data)
	}
}

func TracerouteAPI(w http.ResponseWriter, r *http.Request, data TracerouteInput) {
	result, err := exec.Command("/usr/bin/traceroute", data.Hostname).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s\n", result)
}

func TracerouteForm(w http.ResponseWriter, r *http.Request, data TracerouteInput) {
	out := &TracerouteOutput{}
	out.Hostname = data.Hostname

	resp, err := exec.Command("/usr/bin/traceroute", data.Hostname).Output()
	if err != nil {
		log.Fatal(err)
	}

	// Newlines need to be converted to HTML newlines
	out.Response = strings.Replace(fmt.Sprintf("%s\n", resp), "\n", "<br>", -1)
	result := Results{TracerouteResp: out}
	err = traceTpl.ExecuteTemplate(w, "result.html", result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
