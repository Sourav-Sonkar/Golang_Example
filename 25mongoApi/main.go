package main

import (
	"fmt"
	"github/Sourav-Sonkar/mongoApi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB Api")
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", router.Router()))
	fmt.Println("Listening at 4000")
}
