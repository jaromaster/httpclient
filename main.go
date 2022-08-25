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

	// parse args
	if os.Args[1] == "help" {
		PrintHelp()
		return
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

	silent := false
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--silent") {
			silent = true
			break
		}
	}

	header := ""
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--header=") {
			header = arg[9:]
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
		PrintHelp()
	}

	if len(cookies) > 0 {
		request = SetCookies(request, cookies)
	}
	if len(header) > 0 {
		request = SetHeader(request, header)
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
	fmt.Println(res.Status)
	ct.ResetColor()

	fmt.Println("response time:", elapsed_time, "ms")

	if silent {
		return
	}
	fmt.Println(res.Header)
	fmt.Println(string(resData))
}
