package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("An error has Occured", err)
			return
		}
		defer file.Close()

		content := "Hello , myself vaishnav singh"
		byte, errors := io.WriteString(file, content+"\n")
		fmt.Println("Byte Written is :", byte)
		if errors != nil {
			fmt.Println("Error while creating a file :", errors)
			return
		}

		fmt.Println("successfully created file")
	*/
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("An error has Occured", err)
		return
	}
	defer file.Close()

	// creating a buffer reader for reading files

	buffer := make([]byte, 1024)
	// reading the file in chunks of 1024 bytes
	for {
		n, err := file.Read(buffer)
		if err == io.EOF { //EOF == end of file

			break

		}
		if err != nil {
			fmt.Println("error while reading the file :", err)
			return

		}
		fmt.Println(string(buffer[:n]))
	}

}
