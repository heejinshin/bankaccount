package mydict

import "errors"

type Dictionary map[string]string

var errNotFound := errors.New("Not found")

// Search for a word 
func (d Dictionary) Serach(word string) (string, err) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
	
}



