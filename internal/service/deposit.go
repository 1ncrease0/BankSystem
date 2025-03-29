package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"time"
)

type DepositService struct {
	repo repository.Deposit
}

func NewDepositService(repo repository.Deposit) Deposit {
	return &DepositService{repo: repo}
}

func (s *DepositService) CreateDeposit(dep models.Deposit) (int, error) {
	return s.repo.CreateDeposit(dep)
}

func (s *DepositService) GetDeposit(id int) (models.Deposit, error) {
	return s.repo.Deposit(id)
}

func (s *DepositService) GetDepositsByAccount(accountID int) ([]models.Deposit, error) {
	return s.repo.DepositsByAccount(accountID)
}

func (s *DepositService) DeleteDeposit(id int) error {
	return s.repo.DeleteDeposit(id)
}

func (s *DepositService) GetAllDeposits() ([]models.Deposit, error) {
	return s.repo.GetAllDeposits()
}

func (s *DepositService) UpdateAmount(depositID int, newAmount float64, lastUpdate time.Time) error {
	return s.repo.UpdateAmount(depositID, newAmount, lastUpdate)
}
