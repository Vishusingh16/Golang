package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json: "age"`
	Isadult bool   `json:"isadult"`
}

func main() {
	person := Person{Name: "Vaishnav", Age: 21, Isadult: true}
	fmt.Println("Person data is : ", person)

	//converting person in json (Encoding)format(Marshelling)

	jsonData, err := json.Marshal(person)
	//json.Marshal(person) is used to convert the person object into json format
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println("Person data is ", string(jsonData))

	// Decoding "UnMarshelling
	var personData Person
	err = json.Unmarshal(jsonData, &personData) //& is basically the address
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("person data is :", personData)

}
