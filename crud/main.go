package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}
	data, err := ioutil.ReadAll(res.Body) // it return bytes/slice of data
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data :", string(data))
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding ", err)
		return
	}
	fmt.Println("TODO", todo)
}
func PerformPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Vaishnav Singh",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	//connverting data to string
	jsonstring := string(jsonData)

	jsonReader := strings.NewReader(jsonstring)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

}

func main() {
	fmt.Println("Golearning")
	//	PerformGetRequest()
	PerformPostRequest()

}
