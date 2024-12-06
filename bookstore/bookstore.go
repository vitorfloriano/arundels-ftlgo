package bookstore

import (
	"fmt"
	"errors"
	)

type Book struct {
	Title string
	Author string
	Copies int
	ID int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
	return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return b, nil
}
