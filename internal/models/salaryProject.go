package models

type SalaryProject struct {
	Id                  int     `db:"id"`
	Amount              float64 `db:"amount"`
	ClientAccountId     int     `db:"client_account_id"`
	EnterpriseAccountId int     `db:"enterprise_account_id"`
}

type SalaryProjectRequest struct {
	Id                  int     `db:"id"`
	Amount              float64 `db:"amount"`
	ClientAccountId     int     `db:"client_account_id"`
	EnterpriseAccountId int     `db:"enterprise_account_id"`
	Status              string  `db:"status"`
}
