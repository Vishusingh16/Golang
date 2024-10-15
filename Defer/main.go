package main

import "fmt"

func add(a, b int) int {
	return a + b

}
func main() {
	data := add(5, 6)

	defer fmt.Println("The data is :", data)
	fmt.Println("Dont add all these things")

	//if two defer are added  then they will be executed in reverse order as it start accepingthe defer data in the form of stack

}
