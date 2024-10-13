package main

import "fmt"

func main() {

	studentsGrade := make(map[string]int)
	studentsGrade["John"] = 85
	studentsGrade["vaishnav"] = 94
	studentsGrade["vishu"] = 90
	studentsGrade["Sanskar"] = 9
	studentsGrade["Arav"] = 2
	studentsGrade["Alice"] = 98

	fmt.Println("Marks of bob :", studentsGrade["vaishnav"])
	studentsGrade["vaishnav"] = 100
	fmt.Println("The new  marks of Vaishnav The Great is:", studentsGrade["vaishnav"])
	delete(studentsGrade, "Arav")

	//Check wether key exist in the map or not

	grades, exists := studentsGrade["Arav"]
	fmt.Println("Does Arav exists or not:", exists)
	fmt.Println("Grades of Arav", grades)

}
