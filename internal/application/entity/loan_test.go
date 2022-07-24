package entity

import (
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/golang-hexagonal-architecture/internal/application/entity/enum"
	"github.com/stretchr/testify/assert"
)

func TestLoan(t *testing.T) {
	mockClock := clock.NewMock()
	t.Run("Loan book to a user", func(t *testing.T) {
		loan := defaultLoan(mockClock)
		expectedReturnDate := mockClock.Now().AddDate(0, 0, 15)
		expectedUser := NewUser("email@email.com", "Full name")
		expectedBook := NewBook("A book", "Author name", "Publisher X", "1234567891014")
		expectedStatus := enum.Loaned

		assert.NotNil(t, loan)
		assert.Equal(t, expectedUser, loan.user)
		assert.Equal(t, expectedBook, loan.book)
		assert.Equal(t, expectedReturnDate, loan.returnDate)
		assert.Equal(t, expectedStatus, loan.status)
	})


	t.Run("Return book", func(t *testing.T) {
		loan := defaultLoan(mockClock)
		actualStatus := loan.FinishLoan()
		expectedStatus := enum.Returned
		assert.Equal(t, expectedStatus, actualStatus)
	})


}


func defaultLoan(mockClock *clock.Mock) *Loan {	
	expectedUser := NewUser("email@email.com", "Full name")
	expectedBook := NewBook("A book", "Author name", "Publisher X", "1234567891014")
	return NewLoan(expectedUser, expectedBook, mockClock)
}