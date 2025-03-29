package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"errors"
)

type salaryProjectService struct {
	repo repository.SalaryProject
}

func NewSalaryProjectService(repo repository.SalaryProject) SalaryProject {
	return &salaryProjectService{repo: repo}
}

func (s *salaryProjectService) SalaryProjectByEnterprise(enterpriseId int) ([]models.SalaryProject, error) {
	return s.repo.GetSalaryProjectsByEnterprise(enterpriseId)
}
func (s *salaryProjectService) SalaryProject(id int) (models.SalaryProject, error) {
	return s.repo.GetSalaryProjectById(id)
}

func (s *salaryProjectService) SalaryProjectsByBank(bankId int) ([]models.SalaryProject, error) {
	return s.repo.GetSalaryProjectsByEnterprise(bankId)
}

func (s *salaryProjectService) SalaryProjectsByClient(clientId int) ([]models.SalaryProject, error) {
	return s.repo.GetSalaryProjectsByClient(clientId)
}

func (s *salaryProjectService) CreateSalaryProject(project models.SalaryProject) (int, error) {
	return s.repo.CreateSalaryProject(project)
}

func (s *salaryProjectService) DeleteSalaryProject(id int) error {
	return s.repo.DeleteSalaryProject(id)
}

func (s *salaryProjectService) UpdateSalaryProject(project models.SalaryProject) error {
	return errors.New("UpdateSalaryProject: метод не реализован")
}
