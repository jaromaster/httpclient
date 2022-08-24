package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// syntax:  httpclient method url ["data"] [--save="output-file-path"]

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

	output_file := ""
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--save=") {
			output_file = arg[7:]
		}
	}
	fmt.Println(output_file)

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

	// handle response (print status, response body, ...)
	resData, _ := io.ReadAll(res.Body)

	if len(output_file) > 0 {
		file, err := os.Create(output_file)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(string(resData))
		file.Close()
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)
	fmt.Println(string(resData))
}
