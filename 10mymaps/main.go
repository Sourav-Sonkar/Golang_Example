package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	codingLang := make(map[string]string)

	codingLang["JS"] = "JavaScript"
	codingLang["RB"] = "Ruby"
	codingLang["PY"] = "Python"
	fmt.Println(codingLang)
	fmt.Println("JS is short for:", codingLang["JS"])

	for key, value := range codingLang {
		fmt.Printf("For key %v,value %v\n", key, value)
	}
	//can also use delete for slices
	delete(codingLang, "RB")
	fmt.Println(codingLang)

}
