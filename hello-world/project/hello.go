package main

import "fmt"

func strSliceEmptyOrNil(s []string) bool {
	if len(s) == 0 || s[0] == "" {
		return true
	}
	return false
}

// Greeting is a standard Hello, World greeting
func Greeting(name ...string) string {
	var greetee string
	if strSliceEmptyOrNil(name) {
		greetee = "World"
	} else {
		greetee = name[0]
	}

	greeting := "Hello, " + greetee
	return greeting
}

func main() {
	fmt.Println(Greeting())
}
