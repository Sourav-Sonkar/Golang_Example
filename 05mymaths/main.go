package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcometo maths in golang")

	//var myNum1 int = 4
	//var myNum2 float64 = 6
	//fmt.Println("THe sum is", myNum1+int(myNum2))

	//Random Number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+5)

	//random from crypto

	myrandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myrandomNumber)
}
