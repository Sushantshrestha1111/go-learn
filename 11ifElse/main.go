package main

import "fmt"

func main() {

	fmt.Println("if else in go lang")
	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "less then 10"

	} else if loginCount == 10 {
		result = "more than 10"
	} else {
		result = "number greater than 10"
	}
	fmt.Println(result)

}
