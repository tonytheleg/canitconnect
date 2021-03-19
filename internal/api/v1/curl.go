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

// CurlData stores request data needed to perform a curl
type CurlInput struct {
	URL string `json:"url"`
	//	HTTPProxy  string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// CurlOutput stores response data
type CurlOutput struct {
	URL           string
	Protocol      string
	Status        string
	ContentLength string
	Headers       http.Header
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func Curl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	body, _ := ioutil.ReadAll(r.Body)
	if body != nil {
		fmt.Println("Body is", string(body))
		data := CurlInput{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal json into params")
		}
		fmt.Println("URL is", data.URL)
		CurlAPI(w, r, data.URL)
	}
	CurlForm(w, r, r.FormValue("url"))
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func CurlAPI(w http.ResponseWriter, r *http.Request, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(w, "Something failed", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(w, "Something failed", err)
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "%v %v\nContent Length: %v\n",
		resp.Proto,
		resp.Status,
		resp.ContentLength,
	)
	for k, v := range resp.Header {
		fmt.Fprintf(w, "%v: %v\n", k, v)
	}
}

// Curl takes the passed URL to test and mimics curls response in calling endpoint
func CurlForm(w http.ResponseWriter, r *http.Request, url string) {
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
		Headers:       resp.Header,
	}
	err = tpl.ExecuteTemplate(w, "result.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
