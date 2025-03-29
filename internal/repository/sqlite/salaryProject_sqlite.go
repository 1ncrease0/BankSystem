package sqlite

import (
	"FinanceSystem/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SalaryProjectSQLite struct {
	db *sqlx.DB
}

// NewSalaryProjectSQLite создаёт новый репозиторий для работы с зарплатным проектом
func NewSalaryProjectSQLite(db *sqlx.DB) *SalaryProjectSQLite {

	return &SalaryProjectSQLite{db: db}
}

// CreateSalaryProject добавляет новую запись о зарплатном проекте
func (s *SalaryProjectSQLite) CreateSalaryProject(salaryProject models.SalaryProject) (int, error) {
	query := `
		INSERT INTO salary_project (amount, client_account_id, enterprise_account_id)
		VALUES (?, ?, ?)
	`
	result, err := s.db.Exec(query, salaryProject.Amount, salaryProject.ClientAccountId, salaryProject.EnterpriseAccountId)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании salary_project: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ошибка при получении last insert id: %w", err)
	}
	return int(id), nil
}

// GetSalaryProjectById получает запись зарплатного проекта по её идентификатору
func (s *SalaryProjectSQLite) GetSalaryProjectById(id int) (models.SalaryProject, error) {
	query := `
		SELECT id, amount, client_account_id, enterprise_account_id
		FROM salary_project
		WHERE id = ?
	`
	var salaryProject models.SalaryProject
	err := s.db.Get(&salaryProject, query, id)
	if err != nil {
		return salaryProject, fmt.Errorf("не удалось получить salary_project с id=%d: %w", id, err)
	}
	return salaryProject, nil
}

// GetSalaryProjectsByEnterprise получает все записи зарплатного проекта для указанного счета предприятия
func (s *SalaryProjectSQLite) GetSalaryProjectsByEnterprise(enterpriseAccountId int) ([]models.SalaryProject, error) {
	query := `
		SELECT id, amount, client_account_id, enterprise_account_id
		FROM salary_project
		WHERE enterprise_account_id = ?
	`
	var salaryProjects []models.SalaryProject
	err := s.db.Select(&salaryProjects, query, enterpriseAccountId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить salary_project для enterprise_account_id=%d: %w", enterpriseAccountId, err)
	}
	return salaryProjects, nil
}

// GetSalaryProjectsByClient получает все записи зарплатного проекта для указанного счета клиента
func (s *SalaryProjectSQLite) GetSalaryProjectsByClient(clientAccountId int) ([]models.SalaryProject, error) {
	query := `
		SELECT id, amount, client_account_id, enterprise_account_id
		FROM salary_project
		WHERE client_account_id = ?
	`
	var salaryProjects []models.SalaryProject
	err := s.db.Select(&salaryProjects, query, clientAccountId)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить salary_project для client_account_id=%d: %w", clientAccountId, err)
	}
	return salaryProjects, nil
}

// DeleteSalaryProject удаляет запись зарплатного проекта по её идентификатору
func (s *SalaryProjectSQLite) DeleteSalaryProject(id int) error {
	query := `
		DELETE FROM salary_project
		WHERE id = ?
	`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении salary_project с id=%d: %w", id, err)
	}
	return nil
}
