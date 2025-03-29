package models

import "time"

type Credit struct {
	Id           int        `db:"id"`
	ClientId     int        `db:"client_id"`
	AccountId    int        `db:"account_id"`
	Amount       float64    `db:"amount"`
	Remaining    float64    `db:"remaining"`
	InterestRate float64    `db:"interest_rate"`
	TermMonths   int        `db:"term_months"`
	Status       string     `db:"status"`
	StartDate    *time.Time `db:"start_date"`
	UpdatedAt    time.Time  `db:"updated_at"`
}
