package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter the first number")

	welcom := "welcome to the user"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating :")

	//comma ok || error ok
	// input, error this says you will get input or an error if error store error in error

	//if you dont care about errors then do like down
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating", input)
	fmt.Printf("The type of rating is:%T", input)

}

// func main() {
// 	var num1, num2 int

// 	fmt.Println("enter the first number")
// 	fmt.Scanln(&num1)

// 	fmt.Println("enter second number")
// 	fmt.Scanln(&num2)

// 	result := num1 + num2

// 	fmt.Println(result)

// }
