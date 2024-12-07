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
	PriceCents int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
	return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	b, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return b, nil
}

func NetPriceCents(b Book) int {
	return 0
}
