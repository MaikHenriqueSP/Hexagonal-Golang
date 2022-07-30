package sabao

import "time"

func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}