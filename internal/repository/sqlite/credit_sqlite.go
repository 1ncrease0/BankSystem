package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type CreditSQLite struct {
	db *sqlx.DB
}

func NewCreditSQLite(db *sqlx.DB) *CreditSQLite {
	return &CreditSQLite{db: db}
}

func (r *CreditSQLite) CreateCredit(cr models.Credit) (int, error) {
	query := `INSERT INTO credit 
		(account_id, amount, remaining, interest_rate, term_months,status, start_date, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		cr.AccountId,
		cr.Amount,
		cr.Remaining,
		cr.InterestRate,
		cr.TermMonths,
		cr.Status,
		cr.StartDate,
		cr.UpdatedAt,
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

func (r *CreditSQLite) Credit(id int) (models.Credit, error) {
	var cr models.Credit
	query := `SELECT 
		id, account_id, amount, remaining, interest_rate, term_months, status, start_date, updated_at
		FROM credit WHERE id = ?`
	err := r.db.Get(&cr, query, id)
	return cr, err
}

func (r *CreditSQLite) CreditsByAccount(accountId int) ([]models.Credit, error) {
	var credits []models.Credit
	query := `SELECT 
		id, account_id, amount, remaining, interest_rate, term_months, status, start_date, updated_at
		FROM credit WHERE account_id = ?`
	err := r.db.Select(&credits, query, accountId)
	return credits, err
}

func (r *CreditSQLite) DeleteCredit(id int) error {
	query := `DELETE FROM credit WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CreditSQLite) GetAllCredits() ([]models.Credit, error) {
	var credits []models.Credit
	query := `SELECT 
		id, account_id, amount, remaining, interest_rate, term_months, status, start_date, updated_at
		FROM credit`
	err := r.db.Select(&credits, query)
	return credits, err
}

func (r *CreditSQLite) UpdateRemaining(creditID int, newRemaining float64, updatedAt time.Time) error {
	query := `UPDATE credit SET remaining = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, newRemaining, updatedAt, creditID)
	return err
}
