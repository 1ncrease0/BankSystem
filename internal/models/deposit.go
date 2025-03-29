package models

import "time"

type Deposit struct {
	Id            int        `db:"id"`
	AccountId     int        `db:"account_id"`
	InitialAmount float64    `db:"initial_amount"`
	Amount        float64    `db:"amount"`
	InterestRate  float64    `db:"interest_rate"`
	TermMonths    int        `db:"term_months"`
	Status        string     `db:"status"`
	StartDate     *time.Time `db:"start_date"`
	UpdatedAt     time.Time  `db:"updated_at"`
}
