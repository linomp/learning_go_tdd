package main

import "errors"

type Dictionary map[string]string

var ErrUnknownWord = errors.New("Word not found")

func (d Dictionary) Search(word string) (string, error) {

	// just like indexing array
	definition, keyWasFound := d[word]

	if !keyWasFound {
		return "", ErrUnknownWord
	}

	return definition, nil
}

func Search(dict map[string]string, word string) string {
	return dict[word]
}
