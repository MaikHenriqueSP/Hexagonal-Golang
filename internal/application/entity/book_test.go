package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	t.Run("Book creation", func(t *testing.T){
		expectedTitle := "A Game of thrones"
		expectedAuthor := "George R. R. Martin"
		expectedPublisher := "A Publisher"
		expectedIsbn := "123454567-8"
		book := NewBook(expectedTitle, expectedAuthor, expectedPublisher, expectedIsbn)
		
		assert.NotNil(t, book)
		assert.Equal(t, expectedAuthor, book.author)
		assert.Equal(t, expectedTitle, book.title)
		assert.Equal(t, expectedPublisher, book.publisher)
		assert.Equal(t, expectedIsbn, book.isbn)
	});

	t.Run("Book validation - ISBN length validation", func(t *testing.T) {
		invalidLengthIsbn := "12345"

		expectedTitle := "A Game of thrones"
		expectedAuthor := "George R. R. Martin"
		expectedPublisher := "A Publisher"
		book := NewBook(expectedTitle, expectedAuthor, expectedPublisher, invalidLengthIsbn)

		isBookValid, _ := book.IsValid()

		assert.False(t, isBookValid)
	});

	t.Run("Book validation - ISBN only digits validation", func(t *testing.T) {
		invalidOnlyDigitsIsbn := "123454567891A"

		expectedTitle := "A Game of thrones"
		expectedAuthor := "George R. R. Martin"
		expectedPublisher := "A Publisher"
		book := NewBook(expectedTitle, expectedAuthor, expectedPublisher, invalidOnlyDigitsIsbn)

		isBookValid, _ := book.IsValid()

		assert.False(t, isBookValid)
	});


}
