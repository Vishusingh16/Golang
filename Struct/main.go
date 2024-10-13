package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}
type contact struct {
	Phone string
	Email string
}
type address struct {
	House  int
	Street string
	Area   string
	State  string
}
type employee struct {
	Person_detail person
	Contact_info  contact
	Work_Address  address
}

func main() {
	//var p person
	// p.FirstName = "Vaishnav"
	// p.LastName = "Singh"
	// p.Age = 30

	// person1 := person{
	// 	FirstName: "Vishu",
	// 	FirstName: "Vishu",
	// 	LastName:  "Singh",
	// 	Age:       34,
	// }
	// fmt.Println("Quality of Vishu are as follows:", person1)

	employee1 := employee{
		Person_detail: person{
			FirstName: "Vishu",
			LastName:  "Singh",
			Age:       34,
		},
		Contact_info: contact{
			Phone: "1234567890",
			Email: "vishu@gmail.com",
		},
		Work_Address: address{
			House:  123,
			Street: "Street1",
			Area:   "Area1",
			State:  "State1",
		},
	}
	fmt.Println("The Details of the employees :", employee1)
}
