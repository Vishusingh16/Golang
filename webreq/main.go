package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/")
	if err != nil {
		fmt.Println("Error Getting GET response", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Getting GET response", err)
		return

	}
	fmt.Println("The response is", string(data))

}
