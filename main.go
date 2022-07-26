package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	m := Response{
		Message: "pong",
	}

	br, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(br))
}

func main() {

	http.HandleFunc("/ping", hello)

	http.ListenAndServe(":8090", nil)
}
