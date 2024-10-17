package main

import (
	"fmt"
	"net/url"
)

//Parse :: Converts raw url to an url structure

func main() {
	myURL := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}

	fmt.Println(parseURL.Scheme)
	fmt.Println(parseURL.Path)
	fmt.Println(parseURL.Host)
	fmt.Println(parseURL.RawQuery)

	parseURL.Path = "/newPath"
	fmt.Println(parseURL.Path)
	parseURL.RawQuery = "username-vaishnavsingh"
	fmt.Println(parseURL.RawQuery)

	newUrl := parseURL.String()
	fmt.Println("new url :", newUrl)
}
