package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
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
func SetHeader(req http.Request, header string) http.Request {
	for _, line := range strings.Split(header, `\n`) {
		key_val := strings.SplitN(line, ":", 2)
		req.Header.Add(strings.TrimSpace(key_val[0]), strings.TrimSpace(key_val[1]))
	}

	return req
}

// prepare http request with cookies
func SetCookies(req http.Request, cookies string) http.Request {
	req.Header.Set("Cookie", cookies)
	return req
}

// send http request, return response
func SendRequest(req http.Request) http.Response {

	client := http.Client{}
	res, err := client.Do(&req)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}
