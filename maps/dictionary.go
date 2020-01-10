package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

//Dictionary is a standard map with string as key and value
type Dictionary map[string]string

//ErrNotFound : default error when a word is not found in the dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

//Search : returns definition of word in a dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add adds a new word and definition pair to an existing dictionary
//No pointer needed because a dictionary is a reference type!
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
