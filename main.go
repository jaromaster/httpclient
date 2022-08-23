package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// syntax:  httpclient method url [data]

	// parse args
	if len(os.Args) < 3 {
		log.Fatal("invalid syntax")
	}

	method := os.Args[1]
	url := os.Args[2]

	// TODO setup (cookies, auth, header)

	// execute request
	var res http.Response
	switch method {
	case "get":
		res = Get(url)
	case "post":
		data := os.Args[3]
		res = Post(url, "application/json", data)
	case "put":
		data := os.Args[3]
		res = Put(url, "application/json", data)
	case "delete":
		res = Delete(url)
	default:
		log.Fatal("invalid method")
	}

	// handle response
	data, _ := io.ReadAll(res.Body)

	fmt.Println(string(data))
}
