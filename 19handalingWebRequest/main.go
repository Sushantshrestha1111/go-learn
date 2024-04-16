package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/"

func main() {

	fmt.Println("Handling web Request in go")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of the request is %T\n", response)

	defer response.Body.Close() // it is always caller duty ot close
	data_byte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	contant := string(data_byte)

	fmt.Println(contant)

}
