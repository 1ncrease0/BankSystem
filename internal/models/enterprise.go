package models

type Enterprise struct {
	Id           int    `db:"id"`            // Уникальный идентификатор предприятия
	Type         string `db:"type"`          // Тип предприятия (ИП, ООО, ЗАО и т.д.)
	LegalName    string `db:"legal_name"`    // Юридическое название предприятия
	UNP          string `db:"unp"`           // УНП (уникальный номер предприятия)
	BIC          string `db:"bic"`           // БИК банка предприятия
	LegalAddress string `db:"legal_address"` // Юридический адрес
	BankID       int    `db:"bank_id"`       // Идентификатор банка, в котором зарегистрировано предприятие
}
