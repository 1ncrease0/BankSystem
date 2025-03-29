package sqlite

import (
	"FinanceSystem/internal/models"

	"github.com/jmoiron/sqlx"
	"time"
)

type InstallmentSQLite struct {
	db *sqlx.DB
}

func NewInstallmentSQLite(db *sqlx.DB) *InstallmentSQLite {
	return &InstallmentSQLite{db: db}
}

func (r *InstallmentSQLite) CreateInstallment(inst models.Installment) (int, error) {
	query := `INSERT INTO installment 
		(account_id, amount, remaining, term_months,status, start_date, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?,?)`

	result, err := r.db.Exec(query,
		inst.AccountId,
		inst.Amount,
		inst.Remaining,
		inst.TermMonths,
		inst.Status,
		inst.StartDate,
		inst.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *InstallmentSQLite) Installment(id int) (models.Installment, error) {
	var inst models.Installment
	query := `SELECT 
		id, account_id, amount, remaining, term_months, start_date, updated_at
		FROM installment WHERE id = ?`
	err := r.db.Get(&inst, query, id)
	return inst, err
}

func (r *InstallmentSQLite) InstallmentsByAccount(accountId int) ([]models.Installment, error) {
	var installments []models.Installment
	query := `SELECT 
		id, account_id, amount, remaining, term_months,status, start_date, updated_at
		FROM installment WHERE account_id = ?`
	err := r.db.Select(&installments, query, accountId)
	return installments, err
}

func (r *InstallmentSQLite) DeleteInstallment(id int) error {
	query := `DELETE FROM installment WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *InstallmentSQLite) GetAllInstallments() ([]models.Installment, error) {
	var installments []models.Installment
	query := `SELECT 
		id, account_id, amount, remaining, term_months, status, start_date, updated_at
		FROM installment`
	err := r.db.Select(&installments, query)
	return installments, err
}

func (r *InstallmentSQLite) UpdateRemaining(installmentID int, newRemaining float64, updatedAt time.Time) error {
	query := `UPDATE installment SET remaining = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, newRemaining, updatedAt, installmentID)
	return err
}
