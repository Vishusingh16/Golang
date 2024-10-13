package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple , orange ,banana , prince"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two ka four four two ka one"
	count := strings.Count(str, "four")
	fmt.Println(count)

	str = " Hello Go   "

	trimming := strings.TrimSpace(str)
	fmt.Println(trimming)

	str1 := "Vaishnav"
	str2 := "Singh"
	result := strings.Join([]string{str1, str2}, ".")
	fmt.Println(result)
}
