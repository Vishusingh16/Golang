package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println(i)
	}

	counter := 0
	for {
		fmt.Println("Bhut badi loop")
		counter++
		if counter == 5 {
			break
		}
	}
}
