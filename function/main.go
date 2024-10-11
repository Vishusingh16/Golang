package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("I am learning golang dude")
	ans := add(2, 2)
	fmt.Println(ans)

}
