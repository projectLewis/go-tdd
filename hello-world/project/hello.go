package main

import "fmt"

// Greeting is a standard Hello, World greeting
func Greeting() string {
	greeting := "Hello, World"
	return greeting
}

func main() {
	fmt.Println(Greeting())
}
