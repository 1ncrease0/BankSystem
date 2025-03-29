package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type RegistrationRequestSQLite struct {
	db *sqlx.DB
}

func NewRegistrationRequestSQLite(db *sqlx.DB) *RegistrationRequestSQLite {
	return &RegistrationRequestSQLite{db: db}
}

func (r *RegistrationRequestSQLite) CreateRegistrationRequest(req models.RegistrationRequest) (int, error) {
	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now()
	}

	query := `
		INSERT INTO registration_request (status, client_id, bank_id, created_at)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, req.Status, req.ClientId, req.BankId, req.CreatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *RegistrationRequestSQLite) RegistrationRequestsByBank(bankId int) ([]models.RegistrationRequest, error) {
	var requests []models.RegistrationRequest

	query := `
		SELECT id, status, client_id, bank_id, created_at
		FROM registration_request
		WHERE bank_id = ?
	`
	err := r.db.Select(&requests, query, bankId)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *RegistrationRequestSQLite) UpdateStatus(registrationRequestId int, status string) error {
	query := `
		UPDATE registration_request
		SET status = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query, status, registrationRequestId)
	return err
}
func (r *RegistrationRequestSQLite) DeleteRegistrationRequest(registrationRequestId int) error {
	query := `
		DELETE FROM registration_request
		WHERE id = ?
	`
	_, err := r.db.Exec(query, registrationRequestId)
	return err
}
func (r *RegistrationRequestSQLite) RegistrationRequestById(id int) (models.RegistrationRequest, error) {
	var req models.RegistrationRequest
	err := r.db.QueryRow("SELECT id, client_id, bank_id, status, created_at FROM registration_request WHERE id = ?", id).
		Scan(&req.Id, &req.ClientId, &req.BankId, &req.Status, &req.CreatedAt)
	return req, err
}
