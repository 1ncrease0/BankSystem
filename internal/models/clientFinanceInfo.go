package models

type ClientFinanceInfo struct {
	Accounts     []Account
	Deposits     []Deposit
	Credits      []Credit
	Installments []Installment
}
