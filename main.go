package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// syntax:  httpclient method url ["data"]

	// parse args
	if len(os.Args) < 3 {
		log.Fatal("invalid syntax")
	}
	method := os.Args[1]
	url := os.Args[2]
	data := ""

	if len(os.Args) >= 4 {
		data = os.Args[3]
	}

	// TODO handle cookies, auth, header

	// prepare request
	var request http.Request
	switch method {
	case "get":
		request = PrepareRequest("GET", url, data)
	case "post":
		request = PrepareRequest("POST", url, data)
	case "put":
		request = PrepareRequest("PUT", url, data)
	case "delete":
		request = PrepareRequest("DELETE", url, data)
	default:
		log.Fatal("invalid method")
	}

	// send request
	res := SendRequest(request)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)

	// handle response (print status, response body, ...)
	resData, _ := io.ReadAll(res.Body)

	fmt.Println(string(resData))
}
