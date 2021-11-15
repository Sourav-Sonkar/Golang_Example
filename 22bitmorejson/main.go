package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to json part")
	//EncodeJson()
	DecodeJson()
}

type course struct {
	//alais for json
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              //to hide the field
	Tags     []string `json:"tags,omitempty"` //to hide fields which are empty
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React Js",
		"price": 299,
		"website": "youtube",
		"tags": ["web","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("Json was incorrect")
	}

	//some cases where you want to add data to key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}

func EncodeJson() {
	lcoCourses := []course{
		{
			"React Js",
			299,
			"youtube",
			"bcd123",
			[]string{"web", "js"},
		},
		{
			"Mern",
			199,
			"youtube",
			"abc123",
			[]string{"full stack", "js"},
		}, {
			"Angular",
			399,
			"youtube",
			"789abc",
			nil,
		},
	}
	//package lcocourse data to json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
