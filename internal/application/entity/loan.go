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

func NewLoan(u *User, b *Book, c clock.Clock) (loan *Loan) {
	loanDate := c.Now()
	returnDate := loanDate.AddDate(0, 0, BaseReturnDate)
	return &Loan{u, b, loanDate, returnDate, enum.Loaned, c}
}

func (p *Loan) FinishLoan() {
	if p.status != enum.Loaned {
		return
	}
	
	if p.isLateToReturn() {
		p.status = enum.LateToReturn
		return
	}
	
	p.status = enum.Returned	
}

func (p *Loan) isLateToReturn() (r bool) {
	expectedReturnDate := p.returnDate.AddDate(0, 0, BaseReturnDate)
	return p.clock.Now().AddDate(0, 0, 1).After(expectedReturnDate)
}

