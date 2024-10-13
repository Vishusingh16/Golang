package main

import "fmt"

func main() {
	x := 21
	if x > 20 {
		fmt.Println("x  is greater than 20")

	} else if x < 20 {
		fmt.Println("x is less than 20")
	} else {
		fmt.Println("x is equal to 20")
	}
}
