package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"errors"
)

type SalaryProjectRequestService struct {
	repo           repository.SalaryProjectRequest
	accountService Account
	salaryService  SalaryProject
}

func NewSalaryProjectRequestService(
	repo repository.SalaryProjectRequest,
	accountService Account,
	salaryService SalaryProject,
) SalaryProjectRequest {
	return &SalaryProjectRequestService{
		repo:           repo,
		accountService: accountService,
		salaryService:  salaryService,
	}
}

func (s *SalaryProjectRequestService) CreateSalaryProjectRequest(req models.SalaryProjectRequest) (int, error) {
	// Проверка существования аккаунтов
	if _, err := s.accountService.Account(req.ClientAccountId); err != nil {
		return 0, errors.New("client account not found")
	}
	if _, err := s.accountService.Account(req.EnterpriseAccountId); err != nil {
		return 0, errors.New("enterprise account not found")
	}

	return s.repo.CreateSalaryProjectRequest(req)
}

func (s *SalaryProjectRequestService) GetSalaryProjectRequest(id int) (models.SalaryProjectRequest, error) {
	return s.repo.GetSalaryProjectRequestById(id)
}

func (s *SalaryProjectRequestService) GetSalaryProjectRequestsByClient(clientId int) ([]models.SalaryProjectRequest, error) {
	return s.repo.GetSalaryProjectRequestsByClient(clientId)
}

func (s *SalaryProjectRequestService) GetSalaryProjectRequestsByEnterprise(enterpriseId int) ([]models.SalaryProjectRequest, error) {
	return s.repo.GetSalaryProjectRequestsByEnterprise(enterpriseId)
}

func (s *SalaryProjectRequestService) GetSalaryProjectRequestsByBank(bankId int) ([]models.SalaryProjectRequest, error) {
	return s.repo.GetSalaryProjectRequestsByBank(bankId)
}

func (s *SalaryProjectRequestService) UpdateRequestStatus(id int, status string) error {
	validStatuses := map[string]bool{"pending": true, "approved": true, "rejected": true}
	if !validStatuses[status] {
		return errors.New("invalid status")
	}
	return s.repo.UpdateSalaryProjectRequestStatus(id, status)
}

func (s *SalaryProjectRequestService) DeleteSalaryProjectRequest(id int) error {
	return s.repo.DeleteSalaryProjectRequest(id)
}

func (s *SalaryProjectRequestService) ApproveRequest(requestID int, bankID int) error {
	request, err := s.repo.GetSalaryProjectRequestById(requestID)
	if err != nil {
		return err
	}
	if request.Status != models.RequestUnderConsideration {
		return errors.New("invalid status")
	}

	project := models.SalaryProject{
		ClientAccountId:     request.ClientAccountId,
		EnterpriseAccountId: request.EnterpriseAccountId,
		Amount:              request.Amount,
	}

	if _, err := s.salaryService.CreateSalaryProject(project); err != nil {
		return err
	}

	return s.UpdateRequestStatus(requestID, models.RequestApproved)
}

func (s *SalaryProjectRequestService) RejectRequest(requestID int, bankID int) error {
	r, err := s.repo.GetSalaryProjectRequestById(requestID)
	if err != nil {
		return err
	}
	if r.Status == models.RequestApproved || r.Status == models.RequestRejected {
		return errors.New("заявка уже обработана")
	}
	return s.UpdateRequestStatus(requestID, models.RequestApproved)
}
