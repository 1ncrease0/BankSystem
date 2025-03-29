package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"time"
)

type installmentService struct {
	repo repository.Installment
}

func NewInstallmentService(repo repository.Installment) Installment {
	return &installmentService{repo: repo}
}

func (s *installmentService) CreateInstallment(inst models.Installment) (int, error) {
	return s.repo.CreateInstallment(inst)
}

func (s *installmentService) GetInstallment(id int) (models.Installment, error) {
	return s.repo.Installment(id)
}

func (s *installmentService) GetInstallmentsByAccount(accountID int) ([]models.Installment, error) {
	return s.repo.InstallmentsByAccount(accountID)
}

func (s *installmentService) DeleteInstallment(id int) error {
	return s.repo.DeleteInstallment(id)
}

func (s *installmentService) GetAllInstallments() ([]models.Installment, error) {
	return s.repo.GetAllInstallments()
}

func (s *installmentService) UpdateRemaining(installmentID int, newRemaining float64, updatedAt time.Time) error {
	return s.repo.UpdateRemaining(installmentID, newRemaining, updatedAt)
}
