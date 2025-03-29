package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
)

type FinRequestSQLite struct {
	db *sqlx.DB
}

func NewFinRequestSQLite(db *sqlx.DB) *FinRequestSQLite {
	return &FinRequestSQLite{db: db}
}

func (r *FinRequestSQLite) CreateFinRequest(req models.FinRequest) (int, error) {
	query := `
		INSERT INTO fin_request (
			status, type, client_id, bank_id, account_id, 
			amount, interest_rate, term_months, created_at
		) VALUES (
			:status, :type, :client_id, :bank_id, :account_id, 
			:amount, :interest_rate, :term_months, :created_at
		)`
	result, err := r.db.NamedExec(query, req)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *FinRequestSQLite) FinRequestById(id int) (models.FinRequest, error) {
	var req models.FinRequest
	query := `SELECT * FROM fin_request WHERE id = ?`
	err := r.db.Get(&req, query, id)
	return req, err
}

func (r *FinRequestSQLite) FinRequestsByClient(clientId int) ([]models.FinRequest, error) {
	var requests []models.FinRequest
	query := `SELECT * FROM fin_request WHERE client_id = ?`
	err := r.db.Select(&requests, query, clientId)
	return requests, err
}

func (r *FinRequestSQLite) FinRequestsByBank(bankId int) ([]models.FinRequest, error) {
	var requests []models.FinRequest
	query := `SELECT * FROM fin_request WHERE bank_id = ?`
	err := r.db.Select(&requests, query, bankId)
	return requests, err
}

func (r *FinRequestSQLite) UpdateStatus(finRequestId int, status string) error {
	query := `UPDATE fin_request SET status = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, finRequestId)
	return err
}

func (r *FinRequestSQLite) DeleteFinRequest(finRequestId int) error {
	query := `DELETE FROM fin_request WHERE id = ?`
	_, err := r.db.Exec(query, finRequestId)
	return err
}
