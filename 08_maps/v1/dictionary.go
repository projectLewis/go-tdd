package main

import "errors"

// Dictionary store definitions to words
type Dictionary map[string]string

// ErrNotFound is the error message for missing words in the dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search find a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a key value pair to a dictionary
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
