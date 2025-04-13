package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"errors"
	"time"
)

type AccountService struct {
	repo    repository.Account
	logRepo repository.ActionLog
}

func NewAccountService(repo repository.Account, logRepo repository.ActionLog) *AccountService {
	return &AccountService{
		repo:    repo,
		logRepo: logRepo,
	}
}

func (s *AccountService) TransferMoney(from, to int, amount float64) error {
	fromAccountId, toAccountId := from, to
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	status, err := s.repo.AccountStatus(fromAccountId)
	if err != nil {
		return err
	}
	if status == models.StatusFrozen || status == models.StatusBlocked {
		return errors.New("account is not available")
	}

	aaa, err := s.repo.Account(fromAccountId)
	if err != nil {
		return err
	}
	bankId := aaa.BankId
	
	if err := s.repo.TransferMoney(fromAccountId, toAccountId, amount); err != nil {
		return err
	}

	logEntry := models.ActionLogUnit{
		Type:      models.LogTransfer,
		Time:      time.Now(),
		Sender:    &fromAccountId,
		Recipient: &toAccountId,
		Amount:    &amount,
		BankId:    bankId,
	}
	_, err = s.logRepo.Create(logEntry)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountService) CancelTransfer(logID, fromAccountId, toAccountId int, amount float64) error {

	if err := s.repo.RollbackTransferMoney(fromAccountId, toAccountId, amount); err != nil {
		return err
	}
	
	if err := s.logRepo.Delete(logID); err != nil {
		return err
	}
	return nil
}



func (s *AccountService) PutMoney(accountId int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	status, err := s.repo.AccountStatus(accountId)
	if err != nil {
		return err
	}
	if status == models.StatusFrozen || status == models.StatusBlocked {
		return errors.New("account is not available")
	}
	return s.repo.PlusMoney(accountId, amount)
}

func (s *AccountService) WithdrawMoney(accountId int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	status, err := s.repo.AccountStatus(accountId)
	if err != nil {
		return err
	}
	if status == models.StatusFrozen || status == models.StatusBlocked {
		return errors.New("account is not available")
	}
	return s.repo.MinusMoney(accountId, amount)
}

func (s *AccountService) CreateAccount(account models.Account) (int, error) {
	return s.repo.CreateAccount(account)
}

func (s *AccountService) DeleteAccount(accountId int) error {
	return s.repo.DeleteAccount(accountId)
}

func (s *AccountService) Account(accountId int) (models.Account, error) {
	return s.repo.Account(accountId)
}

func (s *AccountService) AccountsByClient(clientId int) ([]models.Account, error) {
	return s.repo.AccountsByClient(clientId)
}

func (s *AccountService) AccountsByBank(bankId int) ([]models.Account, error) {
	return s.repo.AccountsByBank(bankId)
}

func (s *AccountService) AccountsByEnterprise(enterpriseId int) ([]models.Account, error) {
	return s.repo.AccountsByEnterprise(enterpriseId)
}

func (s *AccountService) AccountsByClientAndBank(clientId, bankId int) ([]models.Account, error) {
	return s.repo.AccountsByClientAndBank(clientId, bankId)
}

func (s *AccountService) Status(accountId int) (string, error) {
	return s.repo.AccountStatus(accountId)
}

func (s *AccountService) ChangeStatus(accountId int, newStatus string) error {
	return s.repo.UpdateStatus(accountId, newStatus)
}
