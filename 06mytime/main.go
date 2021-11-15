package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package of go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2000, time.August, 12, 5, 15, 45, 45, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01/02/2006 Monday 15:04:05"))

}
