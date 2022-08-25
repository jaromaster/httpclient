package main

import "fmt"

// print help
func PrintHelp() {
	usage := "syntax:\t\thttpclient method url [--data='data'] [--save='path'] [--cookies='cookies'] [--silent] [--header='header']"
	methods := "methods:\tGET, POST, PUT, DELETE"
	data := "--data:\t\tsend text or json"
	save := "--save:\t\tspecify path to output file"
	cookies := "--cookies:\tspecify cookies"
	silent := "--silent:\tdo not print response body"
	header := "--header:\tset header values ('key: val\\nkey2: val2')"

	fmt.Println(usage)
	fmt.Println(methods)
	fmt.Println(data)
	fmt.Println(save)
	fmt.Println(cookies)
	fmt.Println(silent)
	fmt.Println(header)
}
