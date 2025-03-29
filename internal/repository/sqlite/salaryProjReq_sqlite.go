package sqlite

import (
	"FinanceSystem/internal/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SalaryProjectRequestSQLite struct {
	db *sqlx.DB
}

func NewSalaryProjectRequestSQLite(db *sqlx.DB) *SalaryProjectRequestSQLite {
	return &SalaryProjectRequestSQLite{db: db}
}

func (r *SalaryProjectRequestSQLite) CreateSalaryProjectRequest(req models.SalaryProjectRequest) (int, error) {
	query := `INSERT INTO salary_project_request 
        (amount, client_account_id, enterprise_account_id, status)
        VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query,
		req.Amount,
		req.ClientAccountId,
		req.EnterpriseAccountId,
		req.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *SalaryProjectRequestSQLite) GetSalaryProjectRequestById(id int) (models.SalaryProjectRequest, error) {
	var req models.SalaryProjectRequest
	query := `SELECT * FROM salary_project_request WHERE id = ?`
	err := r.db.Get(&req, query, id)
	if err == sql.ErrNoRows {
		return req, fmt.Errorf("request not found")
	}
	return req, err
}

func (r *SalaryProjectRequestSQLite) GetSalaryProjectRequestsByBank(bankId int) ([]models.SalaryProjectRequest, error) {
	var requests []models.SalaryProjectRequest
	query := `
		SELECT spr.*
		FROM salary_project_request AS spr
		INNER JOIN account AS a ON spr.enterprise_account_id = a.id
		WHERE a.bank_id = ?`
	err := r.db.Select(&requests, query, bankId)
	fmt.Println(requests, bankId)
	return requests, err
}
func (r *SalaryProjectRequestSQLite) GetSalaryProjectRequestsByClient(clientAccountId int) ([]models.SalaryProjectRequest, error) {
	var requests []models.SalaryProjectRequest
	query := `SELECT * FROM salary_project_request WHERE client_account_id = ?`
	err := r.db.Select(&requests, query, clientAccountId)
	return requests, err
}

func (r *SalaryProjectRequestSQLite) GetSalaryProjectRequestsByEnterprise(enterpriseAccountId int) ([]models.SalaryProjectRequest, error) {
	var requests []models.SalaryProjectRequest
	query := `SELECT * FROM salary_project_request WHERE enterprise_account_id = ?`
	err := r.db.Select(&requests, query, enterpriseAccountId)
	return requests, err
}

func (r *SalaryProjectRequestSQLite) UpdateSalaryProjectRequestStatus(id int, status string) error {
	query := `UPDATE salary_project_request SET status = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *SalaryProjectRequestSQLite) DeleteSalaryProjectRequest(id int) error {
	query := `DELETE FROM salary_project_request WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
