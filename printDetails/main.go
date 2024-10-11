package main

import "fmt"

func main() {
	age := 21
	name := "Vaishnav Singh"
	height := 5.9456789
	fmt.Println("your name :", name, "age is :", age, " height is : ", height)

	fmt.Printf("age is  %d\n", age)
	fmt.Printf("heigh is %.4f\n", height)
	fmt.Printf(" The height is of type  %T\n", height)

}
