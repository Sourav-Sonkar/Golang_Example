package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var responseString strings.Builder

func main() {
	fmt.Println("Welcome to web verb video")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Print("byte count is:", byteCount, "\n")
	fmt.Println(responseString.String())
	//fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"online"	
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	responseString.Write(content)
	fmt.Println(responseString.String())

}

func PerformPostFormRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	//formdata
	data := url.Values{}
	data.Add("First Name", "Sourav")
	data.Add("Last Name", "Sonkar")
	data.Add("Email", "sourav@dev.go")

	respnse, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer respnse.Body.Close()
	content, _ := ioutil.ReadAll(respnse.Body)
	fmt.Println(string(content))
	responseString.Write(content)
	fmt.Println(responseString.String())
	// var myOnlineData map[string]interface{}
	// json.Unmarshal(content, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	// for k, v := range myOnlineData {
	// 	fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	// }
}
