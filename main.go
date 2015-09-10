// Package main provides test api-server hooks
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"encoding/json"
)

var (
	count uint64
)

func main() {
	count = 0

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", apiEcho)
	mux.HandleFunc("/time", apiTime)
	mux.HandleFunc("/count", apiCount)

	http.ListenAndServe(":8080", mux)
}

func apiEcho(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	if body == nil {
		fmt.Fprintf(writer, "")
	} else {
		fmt.Fprintf(writer, "%s", string(body[:]))
	}
}

func apiTime(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Time: %s", time.Now())
}

func apiCount(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "count: %d", count)
	count = count + 1
}
