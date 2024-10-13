package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 6, 7, 8, 5, 33, 563, 211, 3453)
	// fmt.Println("Number :", numbers)
	// fmt.Printf("The number is of data types :%T\n", numbers)
	// fmt.Println("The length ofthe number is :", len(numbers))
	// fmt.Println("The capacity of the slice is :", cap(numbers))

	numbers := make([]int, 2, 4) //(length : 2 , capacity :4)
	numbers = append(numbers, 393)
	fmt.Println("Numbers :", numbers)
	fmt.Println("Numbers :", len(numbers))
	fmt.Println("Numbers :", cap(numbers))
	//if capacity is 2 and there are more elements than 2 in the  slice, it will be reallocated to 4 it just double it everytime

}
