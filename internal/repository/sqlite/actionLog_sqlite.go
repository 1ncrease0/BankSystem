package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
)

type ActionLogSQLite struct {
	db *sqlx.DB
}

func NewActionLogSQLite(db *sqlx.DB) *ActionLogSQLite {
	return &ActionLogSQLite{db: db}
}

func (a *ActionLogSQLite) Create(log models.ActionLogUnit) (int, error) {
	query := `
        INSERT INTO action_log ("type", time, sender, recipient, bank_id, amount)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	res, err := a.db.Exec(query, log.Type, log.Time, log.Sender, log.Recipient, log.BankId, log.Amount)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a *ActionLogSQLite) GetByID(id int) (models.ActionLogUnit, error) {
	var log models.ActionLogUnit
	query := `SELECT * FROM action_log WHERE id = ?`
	err := a.db.Get(&log, query, id)
	return log, err
}

func (a *ActionLogSQLite) GetAll() ([]models.ActionLogUnit, error) {
	var logs []models.ActionLogUnit
	query := `SELECT * FROM action_log`
	err := a.db.Select(&logs, query)
	return logs, err
}

func (a *ActionLogSQLite) GetByBankID(bankId int) ([]models.ActionLogUnit, error) {
	var logs []models.ActionLogUnit
	query := `SELECT * FROM action_log WHERE bank_id = ?`
	err := a.db.Select(&logs, query, bankId)
	return logs, err
}

func (a *ActionLogSQLite) Delete(id int) error {
	query := `DELETE FROM action_log WHERE id = ?`
	_, err := a.db.Exec(query, id)
	return err
}

func (a *ActionLogSQLite) GetByConsumer(consumerId int) ([]models.ActionLogUnit, error) {
	return nil, nil
}
