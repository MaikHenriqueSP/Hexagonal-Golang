package enum

type LoanStatus string


const (
	Loaned LoanStatus = "LOANED"
	Returned LoanStatus = "RETURNED"
	LateToReturn LoanStatus = "LATE_TO_RETURN"
)
