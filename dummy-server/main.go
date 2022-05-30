package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func invocations(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	defer body.Close()

	b, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Fprintf(w, "You've got an error")
	}
	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/invocations", invocations)

	http.ListenAndServe(":8090", nil)
}
