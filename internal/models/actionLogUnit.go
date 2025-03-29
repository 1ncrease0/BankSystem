package models

import "time"

const (
	LogFreeze   = "doFreeze"
	LogBlock    = "doBlock"
	LogTransfer = "doTransfer"
	LogDeposit  = "doDeposit"
	LogWithdraw = "doWithdraw"
)

type ActionLogUnit struct {
	Id        int       `db:"id"`
	Type      string    `db:"type"`
	Time      time.Time `db:"time"`
	Sender    *int      `db:"sender"`
	Recipient *int      `db:"recipient"`
	Amount    *float64  `db:"amount"`
	BankId    int       `db:"bank_id"`
}
