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

func (b *Book) IsValid() (r bool, err error) {
	isIsbnLengthValid := b.isIsbnLengthValid()
	isbnContainsOnlyDigits, err := b.containsOnlyDigits()

	if (err != nil) {
		return r, err
	}
	
	return isIsbnLengthValid && isbnContainsOnlyDigits, err
}

func (b *Book) isIsbnLengthValid() (r bool) {
	return len(b.isbn) == IsbnLength
}

func (b *Book) containsOnlyDigits() (r bool, err error) {	
	return regexp.Match("^[0-9]+$", []byte(b.isbn))
}