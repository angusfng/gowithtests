package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("word not in dictionary") 
	ErrWordExists = errors.New("word already exists in dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil 
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