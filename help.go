package main

import "fmt"

// print help
func PrintHelp() {
	usage := "syntax:\t\thttpclient method url ['data'] [--save='output-file-path'] [--cookies='some-cookies']"
	methods := "methods:\tGET, POST, PUT, DELETE"
	args := "--save: \tspecify path to output file \n--cookies:\tspecify cookies"

	fmt.Println(usage)
	fmt.Println(methods)
	fmt.Println(args)
}
