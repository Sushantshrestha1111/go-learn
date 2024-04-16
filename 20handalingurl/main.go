package main

import (
	"fmt"
	"net/url"
)

const link string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {

	fmt.Print("Handaling url in go")
	fmt.Println(link)

	result, _ := url.Parse(link)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

}
