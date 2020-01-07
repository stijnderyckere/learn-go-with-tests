package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

//Hello : returns a welcome message to the given name
func Hello(name string) string {
	return englishHelloPrefix + name
}
