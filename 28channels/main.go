package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Chanels in golang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChannel	<-	5
	// fmt.Println(<-myChannel)
	wg.Add(2)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch <- 5
		 ch <- 0
		close(ch)
		wg.Done()
	}(myChannel, wg)
	//reciver only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		value, isOpened := <-ch
		fmt.Println(value, isOpened)
		// 	fmt.Println(<-ch)
		//fmt.Println(<-ch)
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
}
