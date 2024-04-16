package main

import "fmt"

func main() {
	var numb = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numb)

	var fruit []string
	fruit = append(fruit, "mango", "banana", "orange")
	fmt.Println("Fruits List:", fruit)
	fruit = append(fruit[1:])
	fmt.Println(fruit)

	//here append method takes thelist and add the items those three values and final product will again be send to fruit

	//how to remove item from slice

	var cources = []string{"java", "js", "css", "python"}
	fmt.Println(cources)
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
