package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to be in a file"

	file, err := os.Create("./mylogs.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./mylogs.txt")

}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)
	fmt.Println("Text data inside the file is\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
