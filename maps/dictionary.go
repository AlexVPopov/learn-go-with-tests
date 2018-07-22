package main

import "errors"

func main() {}

//Search takes a dictionary and a key and returns the value
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// Dictionary is a wrapper around map
type Dictionary map[string]string

// Search takes a key and returns the corresponding value from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
