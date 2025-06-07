package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func test() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

}
func main() {

	// load the enviromnent variables from the .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")

	}

	test()
	port := os.Getenv("PORT")
	// fmt.Println("port from .env file:", port)

	fmt.Println("server started on port:", port)

	http.ListenAndServe("localhost:"+port, nil)

}
