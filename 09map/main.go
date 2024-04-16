package main

import (
	"fmt"
)

func main() {
	fmt.Println("map")
	//make can be used to create map and slice
	//make(map[key]value)
	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println(languages)
	fmt.Println("languages are:", languages["js"])

	//for deleting
	delete(languages, "js")
	fmt.Println(languages)

	//loops
	// using coma ok ignoring key and only taking value
	//:= handles the coma ok for us
	for _, value := range languages {
		fmt.Println("for value is:\n", value)
	}

}
