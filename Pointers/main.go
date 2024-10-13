package main

import "fmt"

func ModifyValue(num *int) {
	*num = *num + 20
}

func main() {

	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num
	// itna likhne se accha h thoda chote me likh lo

	num := "Vaishnav Singh"
	ptr := &num
	//fmt.Println("The value of num  is:", num)
	fmt.Println("The value of ptr is:", ptr)
	fmt.Println("The value inside num is :", num)

	value := 10
	ModifyValue(&value)
	fmt.Println("The value of value is:", value)

}
