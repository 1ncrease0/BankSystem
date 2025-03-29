package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
)

type EnterpriseSQLite struct {
	db *sqlx.DB
}

func NewEnterpriseSQLite(db *sqlx.DB) *EnterpriseSQLite {
	return &EnterpriseSQLite{
		db: db,
	}
}

func (e *EnterpriseSQLite) CreateEnterprise(ent models.Enterprise) (int, error) {
	query := `
		INSERT INTO enterprise 
			(type, legal_name, unp, bic, legal_address, bank_id)
		VALUES 
			(?, ?, ?, ?, ?, ?)
	`
	res, err := e.db.Exec(query, ent.Type, ent.LegalName, ent.UNP, ent.BIC, ent.LegalAddress, ent.BankID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (e *EnterpriseSQLite) EnterpriseById(id int) (models.Enterprise, error) {
	var ent models.Enterprise
	query := `
		SELECT id, type, legal_name, unp, bic, legal_address, bank_id
		FROM enterprise
		WHERE id = ?
	`
	err := e.db.Get(&ent, query, id)
	return ent, err
}

func (e *EnterpriseSQLite) EnterpriseByUNP(unp string) (models.Enterprise, error) {
	var ent models.Enterprise
	query := `
		SELECT id, type, legal_name, unp, bic, legal_address, bank_id
		FROM enterprise
		WHERE unp = ?
	`
	err := e.db.Get(&ent, query, unp)
	return ent, err
}
