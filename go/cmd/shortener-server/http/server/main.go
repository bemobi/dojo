package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	AB string `json:""939393993`
	Ab string
	B  int
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	var data Data
	json.NewDecoder(r.Body).Decode(&data)

	fmt.Fprintf(w, "%v\n", data)
}
