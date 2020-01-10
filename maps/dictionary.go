package main

import (
	"errors"
	"fmt"
)

//Dictionary is a standard map with string as key and value
type Dictionary map[string]string

func main() {
	fmt.Println("Hello world")
	dictionary := Dictionary{"test": "this is just a test"}
	fmt.Println(dictionary)
	dictionary.Add("test2", "this was added via code at runtime")
	fmt.Println(dictionary)
	dictionary.Update("test", "this value was changed via code at runtime")
	fmt.Println(dictionary)
	dictionary.Delete("test")
	fmt.Println(dictionary)
}

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

//Update updates and existing word and definition pair
//No pointer needed because a dictionary is a reference type!
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}

//Delete deletes a word from an existing dictionary
//No pointer needed because a dictionary is a reference type!
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
