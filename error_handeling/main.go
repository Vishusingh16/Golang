package main

import "fmt"

func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("Denominator should not be zero ")
	}
	return a / b, nil

}

func main() {
	fmt.Println("Starting with error handeling")

	ans, error := divide(3, 0)
	//the uderscore  is used as a blank idedntifier , it serves as write only function that allows you to discard values  returned by a function 	or ignore specific values 	 when you are not interested in them'
	if error != nil {
		fmt.Println("error handeling")
	}
	fmt.Println(ans)

}
