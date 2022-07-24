package entity

import "regexp"

type Book struct {
	isbn string
	title string
	author string
	publisher string
}

const IsbnLength int = 13

func NewBook(title, author, publisher, isbn string) *Book {
	return &Book{isbn, title, author, publisher}
}

func (b *Book) IsValid() (bool, error) {
	isIsbnLengthValid := b.isIsbnLengthValid()
	isbnContainsOnlyDigits, err := b.containsOnlyDigits()

	if (err != nil) {
		return false, err
	}
	
	return isIsbnLengthValid && isbnContainsOnlyDigits, nil
}

func (b *Book) isIsbnLengthValid() bool {
	return len(b.isbn) == IsbnLength
}

func (b *Book) containsOnlyDigits() (bool, error) {	
	return regexp.Match("^[0-9]+$", []byte(b.isbn))
}