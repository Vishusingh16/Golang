package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)

	formatted := currentTime.Format("2006/01/02, Monday")
	fmt.Println(formatted)
}
