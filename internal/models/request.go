package models

import "time"

const (
	RequestApproved           = "approved"
	RequestRejected           = "rejected"
	RequestUnderConsideration = "under_consideration"
	RequestTypeCredit         = "credit"
	RequestTypeDeposit        = "deposit"
	RequestTypeInstallment    = "installment"
)

type RegistrationRequest struct {
	Id        int       `db:"id"`
	Status    string    `db:"status"`
	ClientId  int       `db:"client_id"`
	BankId    int       `db:"bank_id"`
	CreatedAt time.Time `db:"created_at"`
}

type FinRequest struct {
	Id           int       `db:"id"`
	Status       string    `db:"status"`
	Type         string    `db:"type"`
	ClientId     int       `db:"client_id"`
	BankId       int       `db:"bank_id"`
	AccountId    int       `db:"account_id"`
	Amount       float64   `db:"amount"`        // Сумма кредита/депозита/рассрочки
	InterestRate float64   `db:"interest_rate"` // Процентная ставка
	TermMonths   int       `db:"term_months"`   // Срок в месяцах
	CreatedAt    time.Time `db:"created_at"`
}
