package entity

import (
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/golang-hexagonal-architecture/internal/application/entity/enum"
	timeUtil "github.com/golang-hexagonal-architecture/pkg/time"
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
		expectedStatus := enum.Returned
		loan.FinishLoan()
		actualStatus := loan.status
		assert.Equal(t, expectedStatus, actualStatus)
	})

	t.Run("Extend loan", func(t *testing.T) {
		loan := defaultLoan(mockClock)
		expectedReturnDate := loan.returnDate.AddDate(0, 0, LoanExtensionDays)
		
		loan.ExtendLoan()

		actualReturnDate := loan.returnDate

		assert.Equal(t, expectedReturnDate, actualReturnDate)
	})

	t.Run("Cannot extend loan if status is different loaned", func(t *testing.T) {
		loan := defaultLoan(mockClock)
		expectedReturnDate := loan.returnDate
		loan.status = enum.Returned
		
		loan.ExtendLoan()

		actualReturnDate := loan.returnDate

		assert.Equal(t, expectedReturnDate, actualReturnDate)
	})

	t.Run("Cannot extend loan if it was already extended three times", func(t *testing.T) {
		loan := defaultLoan(mockClock)
		expectedReturnDate := timeUtil.AddDays(loan.returnDate, LoanExtensionDays * 3)		
		
		loan.ExtendLoan()
		loan.ExtendLoan()
		loan.ExtendLoan()
		loan.ExtendLoan()

		actualReturnDate := loan.returnDate

		assert.Equal(t, expectedReturnDate, actualReturnDate)
	})

}


func defaultLoan(mockClock *clock.Mock) *Loan {	
	expectedUser := NewUser("email@email.com", "Full name")
	expectedBook := NewBook("A book", "Author name", "Publisher X", "1234567891014")
	return NewLoan(expectedUser, expectedBook, mockClock)
}