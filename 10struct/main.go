package main

import "fmt"

func main() {
	sush := User{"sush", "email", 9, true}
	fmt.Println(sush)
	fmt.Printf("value are %v\n", sush)

}

type User struct {
	Name   string
	email  string
	age    int
	status bool
}
