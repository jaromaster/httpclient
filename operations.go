package main

import (
	"bytes"
	"log"
	"net/http"
)

// http GET, return response
func Get(url string) http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}

// http POST, return response
func Post(url string, contentType string, data string) http.Response {
	buffer := bytes.NewBuffer([]byte(data))

	res, err := http.Post(url, contentType, buffer)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}

// http PUT, return response
func Put(url string, contentType string, data string) http.Response {
	buffer := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest(http.MethodPut, url, buffer)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}

// http DELETE, return response
func Delete(url string) http.Response {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return *res
}
