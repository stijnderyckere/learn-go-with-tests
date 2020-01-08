package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("", "world"))
}

//Hello : returns a welcome message to the given name
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	greetingPrefix := englishHelloPrefix

	switch language {
	case "Spanish":
		greetingPrefix = spanishHelloPrefix
	case "French":
		greetingPrefix = frenchHelloPrefix
	}

	return greetingPrefix + name
}
