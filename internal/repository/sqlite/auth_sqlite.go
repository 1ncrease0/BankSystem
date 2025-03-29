package sqlite

import (
	"FinanceSystem/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	bankEmployeeTable         = "bankEmployee"
	enterpriseSpecialistTable = "enterprise_specialist"
	clientsTable              = "client"
)

type AuthSQLite struct {
	db *sqlx.DB
}

func NewAuthSQLite(db *sqlx.DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (r *AuthSQLite) EnterpriseSpecialistById(id int) (models.EnterpriseSpecialist, error) {
	var spec models.EnterpriseSpecialist
	query := `
		SELECT id, username, enterprise_id, password
		FROM enterprise_specialist
		WHERE id = ?
	`
	err := r.db.Get(&spec, query, id)
	return spec, err
}

func (r *AuthSQLite) CreateEnterpriseSpecialist(es models.EnterpriseSpecialist) (int, error) {
	var id int
	query := `
        INSERT INTO enterprise_specialist (
            username, enterprise_id, password
        ) VALUES (?, ?, ?) RETURNING id`
	row := r.db.QueryRow(query, es.UserName, es.EnterpriseId, es.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSQLite) EnterpriseSpecialist(username, password string) (models.EnterpriseSpecialist, error) {
	var es models.EnterpriseSpecialist
	query := fmt.Sprintf("SELECT id, username, enterprise_id, password FROM %s WHERE username=? AND password=?", enterpriseSpecialistTable)
	err := r.db.Get(&es, query, username, password)
	return es, err
}
func (r *AuthSQLite) CreateClient(client models.Client) (int, error) {
	var id int
	query := `
        INSERT INTO client (
            name, surname, patronymic, username, 
            passport_series, passport_number, id_number, 
            phone_number, email, password
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id`

	row := r.db.QueryRow(query,
		client.Name, client.Surname, client.Patronymic, client.UserName,
		client.PassportSeries, client.PassportNumber, client.IdNumber,
		client.PhoneNumber, client.Email, client.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSQLite) ClientById(clientId int) (models.Client, error) {
	var client models.Client
	query := fmt.Sprintf("SELECT id,  name, surname, patronymic, username, passport_series, passport_number, id_number, phone_number, email, password FROM %s WHERE id=?", clientsTable)
	err := r.db.Get(&client, query, clientId)
	return client, err
}

func (r *AuthSQLite) Client(username, password string) (models.Client, error) {
	var client models.Client
	query := fmt.Sprintf("SELECT id,  name, surname, patronymic, username, passport_series, passport_number, id_number, phone_number, email, password FROM %s WHERE username=? AND password=?", clientsTable)
	err := r.db.Get(&client, query, username, password)
	return client, err
}

func (r *AuthSQLite) BankEmployee(username, password string) (models.BankEmployee, error) {
	var employee models.BankEmployee
	query := fmt.Sprintf("SELECT id,username, role, bank_id, password FROM %s WHERE username=? AND password=?", bankEmployeeTable)
	err := r.db.Get(&employee, query, username, password)
	return employee, err
}

func (r *AuthSQLite) CreateBankEmployee(employee models.BankEmployee) (int, error) {
	var id int
	query := `
        INSERT INTO bankEmployee(
           username, role, bank_id, password
        ) VALUES ( ?, ?, ?, ?) RETURNING id`

	row := r.db.QueryRow(query,
		employee.UserName, employee.Role, employee.BankId, employee.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
