package main

import "fmt"

func main() {
	fmt.Sprintln("Welcome to the loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	/* can also use comma ok syntax*/
	// for index, day := range days {
	// 	fmt.Printf("Index is %v,value is %v\n", index, day)
	// }

	rougeValue := 0
	/* example of while loop*/
	for rougeValue < 10 {
		//example of goto statement
		if rougeValue == 2 {
			goto gotoStatement
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		if rougeValue == 8 {
			break
		}
		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}

gotoStatement:
	fmt.Println("THis the goto statement")
}
