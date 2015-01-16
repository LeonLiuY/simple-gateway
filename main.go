package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	r.RequestURI = ""
	r.Host = "httpbin.org"
	r.URL.Scheme = "http"
	r.URL.Host = "httpbin.org"
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	resp.Write(w)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
