package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usual are pointers
var mut sync.Mutex //usual are pointers

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")
	websites := []string{"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.fb.com",
		"https://www.github.com",
	}
	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Some error in here")
	} else {
		//lock the memory
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
