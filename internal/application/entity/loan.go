package entity

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang-hexagonal-architecture/internal/application/entity/enum"
)

type Loan struct {
	user *User
	book *Book
	loanDate time.Time
	returnDate time.Time
	status enum.LoanStatus
	clock clock.Clock
}

const BaseReturnDate int = 15

func NewLoan(user *User, book *Book, clock clock.Clock) (loan *Loan) {
	loanDate := clock.Now()
	returnDate := loanDate.AddDate(0, 0, BaseReturnDate)
	return &Loan{user, book, loanDate, returnDate, enum.Loaned, clock}
}

func (p *Loan) FinishLoan() enum.LoanStatus {
	if p.status != enum.Loaned {
		return p.status
	}
	
	if p.isLateToReturn() {
		p.status = enum.LateToReturn
		return enum.LateToReturn
	}
	
	p.status = enum.Returned
	return enum.Returned
}

func (p *Loan) isLateToReturn() bool {
	expectedReturnDate := p.returnDate.AddDate(0, 0, BaseReturnDate)
	return p.clock.Now().AddDate(0, 0, 1).After(expectedReturnDate)
}

