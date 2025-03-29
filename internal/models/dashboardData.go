package models

type DashboardData struct {
	Accounts     []Account
	Credits      []Credit
	Deposits     []Deposit
	Installments []Installment
	Client       Client
	Bank         Bank
}
