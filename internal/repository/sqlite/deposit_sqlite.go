package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type DepositSQLite struct {
	db *sqlx.DB
}

func NewDepositSQLite(db *sqlx.DB) *DepositSQLite {
	return &DepositSQLite{db: db}
}

func (r *DepositSQLite) GetAllDeposits() ([]models.Deposit, error) {
	var deposits []models.Deposit
	query := `SELECT 
		id, account_id, initial_amount, amount, interest_rate, term_months, status, start_date, updated_at
		FROM deposit`
	err := r.db.Select(&deposits, query)
	return deposits, err
}

func (r *DepositSQLite) UpdateAmount(depositID int, newAmount float64, lastUpdate time.Time) error {
	query := `UPDATE deposit SET amount = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, newAmount, lastUpdate, depositID)
	return err
}

func (r *DepositSQLite) CreateDeposit(dep models.Deposit) (int, error) {
	query := `INSERT INTO deposit 
		(account_id, initial_amount, amount, interest_rate, term_months, status,start_date, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		dep.AccountId,
		dep.InitialAmount,
		dep.Amount,
		dep.InterestRate,
		dep.TermMonths,
		dep.Status,
		dep.StartDate,
		dep.UpdatedAt,
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

func (r *DepositSQLite) Deposit(id int) (models.Deposit, error) {
	var dep models.Deposit
	query := `SELECT 
		id, account_id, initial_amount, amount, interest_rate, term_months,status, start_date, updated_at
		FROM deposit WHERE id = ?`
	err := r.db.Get(&dep, query, id)
	return dep, err
}

func (r *DepositSQLite) DepositsByAccount(accountId int) ([]models.Deposit, error) {
	var deposits []models.Deposit
	query := `SELECT 
		id, account_id, initial_amount, amount, interest_rate, term_months, status, start_date, updated_at
		FROM deposit WHERE account_id = ?`
	err := r.db.Select(&deposits, query, accountId)
	return deposits, err
}

func (r *DepositSQLite) DeleteDeposit(id int) error {
	query := `DELETE FROM deposit WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
