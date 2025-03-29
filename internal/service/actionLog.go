package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
)

type actionLogService struct {
	repo repository.ActionLog
}

func NewActionLogService(repo repository.ActionLog) ActionLog {
	return &actionLogService{repo: repo}
}

func (s *actionLogService) Create(log models.ActionLogUnit) (int, error) {

	return s.repo.Create(log)
}

func (s *actionLogService) GetByID(id int) (models.ActionLogUnit, error) {
	return s.repo.GetByID(id)
}

func (s *actionLogService) GetAll() ([]models.ActionLogUnit, error) {
	return s.repo.GetAll()
}

func (s *actionLogService) GetByBankID(bankId int) ([]models.ActionLogUnit, error) {
	return s.repo.GetByBankID(bankId)
}

func (s *actionLogService) Delete(id int) error {
	return s.repo.Delete(id)
}
