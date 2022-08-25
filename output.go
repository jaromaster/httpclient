package main

import (
	"fmt"
	"net/http"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
)

// print status
func PrintStatus(code int, status string) {
	fmt.Print("status: ")
	if code >= 200 && code < 300 {
		ct.Foreground(ct.Green, false)
	} else if code >= 300 && code < 400 {
		ct.Foreground(ct.Yellow, false)
	} else {
		ct.Foreground(ct.Red, false)
	}
	fmt.Println(status)
	ct.ResetColor()
}

// print response time
func PrintResponseTime(t int64) {
	fmt.Println("response time:", t, "ms")
}

// print header
func PrintHeader(header http.Header) {
	fmt.Println("Header:")
	for k, v := range header {
		fmt.Printf("%s:  %s\n", k, strings.Join(v, " "))
	}
}

// print response body
func PrintData(data string) {
	fmt.Println("Body:")
	fmt.Println(data)
}
