package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrKeyAlreadyExists = errors.New("Tried to add value to key that already exists")
var ErrKeyDoesNotExist = errors.New("Key does not exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
		return nil
	}
	return ErrKeyAlreadyExists
}

func (d Dictionary) Update(key, value string) error {
	_, ok := d[key]
	if ok {
		d[key] = value
		return nil
	}
	return ErrKeyDoesNotExist
}
func (d Dictionary) Delete(key string) error {
	_, ok := d[key]
	if ok {
		delete(d, key)
		return nil
	}
	return ErrKeyDoesNotExist

}
