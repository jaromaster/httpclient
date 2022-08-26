package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// parse args
	if os.Args[1] == "help" {
		PrintHelp()
		return
	}

	// always set
	method := os.Args[1]
	url := os.Args[2]

	// optional flags
	argsMap := CreateFlagsFromArgs(os.Args)
	data := argsMap["--data"]
	output_file := argsMap["--save"]
	cookies := argsMap["--cookies"]
	silent := len(argsMap["--silent"]) > 0
	header := argsMap["--header"]
	auth := argsMap["--auth"]
	user_agent := argsMap["--useragent"]

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
		PrintHelp()
		return
	}

	if len(cookies) > 0 {
		request = SetCookies(request, cookies)
	}
	if len(header) > 0 {
		request = SetHeader(request, header)
	}
	if len(auth) > 0 {
		request = SetAuth(request, auth)
	}
	if len(user_agent) > 0 {
		request = SetUserAgent(request, user_agent)
	}

	// send request and measure response time
	start_time := time.Now().UnixMilli()
	res := SendRequest(request)
	resData, _ := io.ReadAll(res.Body)
	elapsed_time := time.Now().UnixMilli() - start_time

	// write to file
	if len(output_file) > 0 {
		file, err := os.Create(output_file)
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString(string(resData))
		file.Close()
	}

	// status and time
	PrintStatus(res.StatusCode, res.Status)
	PrintResponseTime(elapsed_time)

	if silent {
		return
	}
	// response header
	fmt.Println()
	PrintHeader(res.Header)

	// response body
	fmt.Println()
	PrintData(string(resData))
}
