package main

import "fmt"

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(GoodBye())
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	return "hello world"
}

func GoodBye() string {
	return "Thanks for coming"
}
