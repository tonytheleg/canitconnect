package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Curl(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	body, _ := ioutil.ReadAll(r.Body)
	params := make(map[string]string)
	err := json.Unmarshal(body, &params)
	if err != nil {
		log.Println("Failed to unmarshal json into params")
	}

	//userId, title, id := params["userId"], params["title"], params["id"]
	json_data, err := json.Marshal(params)
	if err != nil {
		log.Println("Failed to create json_data")
	}
	//fmt.Fprintf(w, "User ID: %s\nTitle: %s\nID: %s\n", userId, title, id)
	fmt.Fprintln(w, "Method is", r.Method)
	fmt.Fprintln(w, "Json Data is", string(json_data[:]))
	if r.Method == "GET" {
		resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Fprintln(w, "Body:\n", string(body[:]))
	}
	if r.Method == "POST" {
		resp, err := client.Post(
			"https://jsonplaceholder.typicode.com/posts",
			"application/json",
			bytes.NewBuffer(json_data),
		)

		if err != nil {
			log.Fatal(err)
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]string)
		err = json.Unmarshal(respBody, &data)
		if err != nil {
			log.Println("Failed to unmarshal json into data")
		}

		//userId, title, id := params["userId"], params["title"], params["id"]
		out, err := json.Marshal(data)
		if err != nil {
			log.Println("Failed to create json_data")
		}
		fmt.Fprintf(w, "Response JSON: %v", string(out[:]))
		/*fmt.Fprintln(w, "Response Body:", resp.Body)
		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		fmt.Fprintln(w, "Response JSON", res)
		*/
	}

	// parse data
	//body, _ := ioutil.ReadAll(r.Body)
	//params := make(map[string]string)
	//_ = json.Unmarshal(body, &params)

}
