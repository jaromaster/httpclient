package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	ct "github.com/daviddengcn/go-colortext"
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
			break
		}
	}

	cookies := ""
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--cookies=") {
			cookies = arg[10:]
			break
		}
	}

	// TODO handle auth, header

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

	if len(cookies) > 0 {
		request = SetCookies(request, cookies)
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

	// color status
	fmt.Print("status: ")
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		ct.Foreground(ct.Green, false)
	} else if res.StatusCode >= 300 && res.StatusCode < 400 {
		ct.Foreground(ct.Yellow, false)
	} else {
		ct.Foreground(ct.Red, false)
	}
	fmt.Println(res.StatusCode)
	ct.ResetColor()

	// output
	fmt.Println("response time:", elapsed_time, "ms")
	fmt.Println(res.Header)
	fmt.Println(string(resData))
}
