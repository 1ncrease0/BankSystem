package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"time"
)

type CreditService struct {
	repo repository.Credit
}

func NewCreditService(repo repository.Credit) Credit {
	return &CreditService{repo: repo}
}

func (s *CreditService) CreateCredit(cr models.Credit) (int, error) {
	return s.repo.CreateCredit(cr)
}

func (s *CreditService) GetCredit(id int) (models.Credit, error) {
	return s.repo.Credit(id)
}

func (s *CreditService) GetCreditsByAccount(accountID int) ([]models.Credit, error) {
	return s.repo.CreditsByAccount(accountID)
}

func (s *CreditService) DeleteCredit(id int) error {
	return s.repo.DeleteCredit(id)
}

func (s *CreditService) GetAllCredits() ([]models.Credit, error) {
	return s.repo.GetAllCredits()
}

func (s *CreditService) UpdateRemaining(creditID int, newRemaining float64, updatedAt time.Time) error {
	return s.repo.UpdateRemaining(creditID, newRemaining, updatedAt)
}
