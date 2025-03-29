package models

import "time"

const (
	StatusBlocked            = "blocked"
	StatusFrozen             = "frozen"
	StatusAvailable          = "available"
	StatusUnderConsideration = "underConsideration"
)

type Account struct {
	Id           int       `json:"id"             db:"id"`
	ClientId     *int      `json:"client_id"      db:"client_id"`
	EnterpriseId *int      `json:"enterprise_id"  db:"enterprise_id"`
	BankId       int       `json:"bank_id"        db:"bank_id"`
	Currency     string    `json:"currency"       db:"currency"`
	Balance      float64   `json:"balance"        db:"balance"`
	Status       string    `json:"status"         db:"status"`
	LastUpdate   time.Time `json:"last_update"    db:"last_update"`
}
