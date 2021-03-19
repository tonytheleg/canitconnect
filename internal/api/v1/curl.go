package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

// CurlData stores info needed to perform a curl
type CurlInput struct {
	URL string `json:"url"`
	//	HTTPProxy  string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

type CurlOutput struct {
	URL           string
	Protocol      string
	Status        string
	ContentLength string
	ContentType   string
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func Curl(w http.ResponseWriter, r *http.Request) {
	// Get data from original request to form the curl command
	body, _ := ioutil.ReadAll(r.Body)
	data := CurlInput{}
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
	fmt.Fprintf(w, "%v %v\ncontent-length: %v\ncontent-type: %v\n",
		resp.Proto,
		resp.Status,
		resp.ContentLength,
		resp.Header["Content-Type"][0],
	)
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func CurlForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	url := r.FormValue("url")

	// Create the new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(w, "Something failed", err)
	}
	defer resp.Body.Close()
	data := CurlOutput{
		URL:           url,
		Protocol:      resp.Proto,
		Status:        resp.Status,
		ContentLength: fmt.Sprint(resp.ContentLength),
		ContentType:   resp.Header["Content-Type"][0],
	}
	err = tpl.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
