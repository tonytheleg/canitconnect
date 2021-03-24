package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var curltpl *template.Template

func init() {
	curltpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

// CurlInput stores request data needed to perform a curl
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

// Curl confirms the request source and directs to the correct curl method
func Curl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.URL.String() == "/api/v1/curl" {
		body, _ := ioutil.ReadAll(r.Body)
		data := CurlInput{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			fmt.Fprintln(w, "Failed to unmarshal json into params")
		}
		CurlAPI(w, r, data.URL)
	} else {
		CurlForm(w, r, r.FormValue("url"))
	}
}

// CurlAPI takes the direct passed data to test and mimics curls' response in calling the endpoint
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

// CurlForm parses the app form to test and mimics curls' response in calling the endpoint
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
	data := &CurlOutput{
		URL:           url,
		Protocol:      resp.Proto,
		Status:        resp.Status,
		ContentLength: fmt.Sprint(resp.ContentLength),
		Headers:       resp.Header,
	}
	result := Results{CurlResp: data}
	err = curltpl.ExecuteTemplate(w, "result.html", result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
