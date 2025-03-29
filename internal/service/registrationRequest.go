package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"errors"
	"sort"
	"time"
)

type registrationRequestService struct {
	repo           repository.RegistrationRequest
	accountService Account
}

func NewRegistrationRequestService(repo repository.RegistrationRequest, accountService Account) RegistrationRequest {
	return &registrationRequestService{
		repo:           repo,
		accountService: accountService,
	}
}

func (s *registrationRequestService) CreateRegistrationRequest(reg models.RegistrationRequest) (int, error) {
	return s.repo.CreateRegistrationRequest(reg)
}

func (s *registrationRequestService) RegistrationRequestsByBank(bankId int) ([]models.RegistrationRequest, error) {
	r, err := s.repo.RegistrationRequestsByBank(bankId)
	if err != nil {
		return nil, err
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].CreatedAt.After(r[j].CreatedAt)
	})
	return r, nil
}

func (s *registrationRequestService) UpdateStatus(registrationRequestId int, status string) error {
	return s.repo.UpdateStatus(registrationRequestId, status)
}

func (s *registrationRequestService) DeleteRegistrationRequest(registrationRequestId int) error {

	return s.repo.DeleteRegistrationRequest(registrationRequestId)
}

func (s *registrationRequestService) ApproveRequest(reqId, bankId int) error {

	req, err := s.repo.RegistrationRequestById(reqId)
	if err != nil {
		return err
	}

	if req.BankId != bankId {
		return err
	}
	if req.Status == models.RequestRejected || req.Status == models.RequestApproved {
		return errors.New("заявка уже обработана")
	}

	accounts, err := s.accountService.AccountsByClientAndBank(req.ClientId, bankId)
	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		newAccount := models.Account{
			ClientId:     &req.ClientId,
			EnterpriseId: nil,
			BankId:       bankId,
			Currency:     "RUB",
			Balance:      0.0,
			Status:       models.StatusAvailable,
			LastUpdate:   time.Now(),
		}
		_, err := s.accountService.CreateAccount(newAccount)
		if err != nil {
			return err
		}
	}

	return s.repo.UpdateStatus(reqId, models.RequestApproved)
}

func (s *registrationRequestService) RejectRequest(reqId, bankId int) error {
	r, err := s.repo.RegistrationRequestById(reqId)
	if err != nil {
		return err
	}
	if r.Status == models.RequestApproved || r.Status == models.RequestRejected {
		return errors.New("заявка уже обработана")
	}
	return s.repo.UpdateStatus(reqId, models.RequestRejected)
}
