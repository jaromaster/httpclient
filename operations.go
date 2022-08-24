package main

import (
	"bytes"
	"log"
	"net/http"
)

// prepare http request
func PrepareRequest(method string, url string, data string) http.Request {
	buffer := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		log.Fatal(err)
	}

	return *req
}

// prepare http request with header

// prepare http request with cookies

// send http request, return response
func SendRequest(req http.Request) http.Response {

	client := http.Client{}
	res, err := client.Do(&req)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}
