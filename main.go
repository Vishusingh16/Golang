package main

import (
	"fmt"
	"mylearnings/myutils"
)

func main() {
	fmt.Println("Hello world")

	myutils.PrintMessage("Hellow World")

	var name = "Vaishnav Singh"
	fmt.Println(name)

	var currency = 67000
	fmt.Println(currency)

	var dimension float64 = 32.34
	fmt.Println(dimension)

	var decesion bool = false
	fmt.Println(decesion)

	person := " vishu singh"
	fmt.Println(person)

}

func PublicFunction() {
	fmt.Println("this function can be accessed in other packages , classes etc")
}

func privateFunctions() {

	fmt.Println(" This function can be accessed in only  this package, this class and nowhere other")

}
