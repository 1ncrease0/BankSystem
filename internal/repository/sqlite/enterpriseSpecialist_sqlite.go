package sqlite

import (
	"FinanceSystem/internal/models"
	"github.com/jmoiron/sqlx"
)

type EnterpriseSpecialistSQLite struct {
	db *sqlx.DB
}

func NewEnterpriseSpecialistSQLite(db *sqlx.DB) *EnterpriseSpecialistSQLite {
	return &EnterpriseSpecialistSQLite{
		db: db,
	}
}

func (es *EnterpriseSpecialistSQLite) CreateEnterpriseSpecialist(spec models.EnterpriseSpecialist) (int, error) {
	query := `
		INSERT INTO enterprise_specialist 
			(username, enterprise_id, password)
		VALUES 
			(?, ?, ?)
	`
	res, err := es.db.Exec(query, spec.UserName, spec.EnterpriseId, spec.Password)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (es *EnterpriseSpecialistSQLite) EnterpriseSpecialist(username, password string) (models.EnterpriseSpecialist, error) {
	var spec models.EnterpriseSpecialist
	query := `
		SELECT id, username, enterprise_id, password
		FROM enterprise_specialist
		WHERE username = ? AND password = ?
	`
	err := es.db.Get(&spec, query, username, password)
	return spec, err
}

func (es *EnterpriseSpecialistSQLite) EnterpriseSpecialistById(id int) (models.EnterpriseSpecialist, error) {
	var spec models.EnterpriseSpecialist
	query := `
		SELECT id, username, enterprise_id, password
		FROM enterprise_specialist
		WHERE id = ?
	`
	err := es.db.Get(&spec, query, id)
	return spec, err
}
