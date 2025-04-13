package models

type Enterprise struct {
	Id           int    `db:"id"`          
	Type         string `db:"type"`          
	LegalName    string `db:"legal_name"`   
	UNP          string `db:"unp"`          
	BIC          string `db:"bic"`          
	LegalAddress string `db:"legal_address"`
	BankID       int    `db:"bank_id"`       
}
