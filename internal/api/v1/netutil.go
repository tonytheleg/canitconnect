package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ShowCurlData is just a test to see how the data passed by curl
func ShowCurlData(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	params := make(map[string]string)
	_ = json.Unmarshal(body, &params)

	fmt.Fprintf(w, "Params: %v\n", params)
}
