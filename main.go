package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to our todo application!")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/get-name", printName)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User"
	fmt.Fprintln(writer, greeting)
}

func printName(resp http.ResponseWriter, req *http.Request) {
	var name string = "Adnan"
	fmt.Fprintln(resp, name)
}
