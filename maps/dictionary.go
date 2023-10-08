package main

import "errors"

func Search(dictionary map[string]string, word string) string {

	return dictionary[word]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add the word because it already exists")
)