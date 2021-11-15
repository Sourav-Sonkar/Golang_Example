package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn/?coursename=reactjs&paymentid=ddjdjj"

func main() {
	fmt.Println("Welcome to handling of URLSs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	qparms := result.Query()

	fmt.Printf("The query parameter type is %T\n", qparms)
	fmt.Println(qparms["coursename"])

	for index, val := range qparms {
		fmt.Println("Param and value is:", index, val)
	}

	//always pass reference of the url
	partsOFUrl := &url.URL{
		Scheme:   result.Scheme,
		Host:     result.Host,
		RawQuery: result.RawQuery,
		Path:     result.Path,
	}
	anotherURL := partsOFUrl.String()
	fmt.Println(anotherURL)
}
