package v1

import "net/http"

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

// NetcatInput stores info needed to perform a netcat
type NetcatInput struct {
	Hostname string `json:"host"`
	Port     string `json:"port"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// NetcatOutput contains the data for a traceroute response
type NetcatOutput struct {
	Hostname string
	Response string
}

// TracerouteInput stores info needed to perform a Traceroute
type TracerouteInput struct {
	Hostname string `json:"host"`
	//  HTTPProxy string `json:"http_proxy"`
	//  HTTPSProxy string `json:"https_proxy"`
}

// TracerouteOutput contains the data for a traceroute response
type TracerouteOutput struct {
	Hostname string
	Response string
}

type Results struct {
	CurlResp       *CurlOutput
	NetcatResp     *NetcatOutput
	TracerouteResp *TracerouteOutput
}
