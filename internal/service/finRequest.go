package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"errors"
	"sort"
)

type finRequestService struct {
	repo               repository.FinRequest
	accountService     Account
	depositService     Deposit
	creditService      Credit
	installmentService Installment
}

func NewFinRequestService(repo repository.FinRequest, account Account, deposit Deposit, credit Credit, installment Installment) FinRequest {
	return &finRequestService{
		repo:               repo,
		accountService:     account,
		depositService:     deposit,
		creditService:      credit,
		installmentService: installment,
	}
}

func (s *finRequestService) CreateFinRequest(req models.FinRequest) (int, error) {
	acc, err := s.accountService.AccountsByClientAndBank(req.ClientId, req.BankId)
	if err != nil {
		return 0, err
	}
	is := false
	id := 0
	for _, v := range acc {
		if v.Id == req.AccountId {
			is = true
			id = v.Id
			break
		}
	}

	if !is {
		return 0, errors.New("счет недоступен")
	}
	status, err := s.accountService.Status(id)
	if err != nil {
		return 0, err
	}
	if status != models.StatusAvailable {
		return 0, errors.New("счет недоступен")
	}
	return s.repo.CreateFinRequest(req)
}

func (s *finRequestService) GetFinRequestByID(id int) (models.FinRequest, error) {
	return s.repo.FinRequestById(id)
}

func (s *finRequestService) GetFinRequestsByClient(clientID int) ([]models.FinRequest, error) {

	return s.repo.FinRequestsByClient(clientID)
}

func (s *finRequestService) GetFinRequestsByBank(bankID int) ([]models.FinRequest, error) {
	r, err := s.repo.FinRequestsByBank(bankID)
	if err != nil {
		return nil, err
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].CreatedAt.After(r[j].CreatedAt)
	})
	return r, nil
}

func (s *finRequestService) ChangeStatus(finRequestID int, status string) error {
	return s.repo.UpdateStatus(finRequestID, status)
}

func (s *finRequestService) RemoveFinRequest(finRequestID int) error {
	return s.repo.DeleteFinRequest(finRequestID)
}

func (s *finRequestService) ApproveRequest(finReq models.FinRequest) error {
	currentReq, err := s.repo.FinRequestById(finReq.Id)
	if err != nil {
		return err
	}
	if currentReq.Status == models.RequestApproved || currentReq.Status == models.RequestRejected {
		return errors.New("заявка уже обработана")
	}

	if err := s.repo.UpdateStatus(finReq.Id, models.RequestApproved); err != nil {
		return err
	}

	//now := time.Now()
	switch finReq.Type {
	case models.RequestTypeCredit:
		credit := models.Credit{
			ClientId:     finReq.ClientId,
			AccountId:    finReq.AccountId,
			Amount:       finReq.Amount,
			Remaining:    finReq.Amount,
			InterestRate: finReq.InterestRate,
			TermMonths:   finReq.TermMonths,
			Status:       models.StatusAvailable,
			StartDate:    &finReq.CreatedAt,
			UpdatedAt:    finReq.CreatedAt,
		}
		if _, err := s.creditService.CreateCredit(credit); err != nil {
			return err
		}
	case models.RequestTypeDeposit:
		deposit := models.Deposit{
			AccountId:     finReq.AccountId,
			InitialAmount: finReq.Amount,
			Amount:        finReq.Amount,
			InterestRate:  finReq.InterestRate,
			TermMonths:    finReq.TermMonths,
			Status:        models.StatusAvailable,
			StartDate:     &finReq.CreatedAt,
			UpdatedAt:     finReq.CreatedAt,
		}
		if _, err := s.depositService.CreateDeposit(deposit); err != nil {
			return err
		}
	case models.RequestTypeInstallment:
		installment := models.Installment{
			ClientId:   finReq.ClientId,
			AccountId:  finReq.AccountId,
			Amount:     finReq.Amount,
			Remaining:  finReq.Amount,
			TermMonths: finReq.TermMonths,
			Status:     models.StatusAvailable,
			StartDate:  &finReq.CreatedAt,
			UpdatedAt:  finReq.CreatedAt,
		}
		if _, err := s.installmentService.CreateInstallment(installment); err != nil {
			return err
		}
	default:
		return errors.New("неизвестный тип запроса")
	}

	return nil
}

func (s *finRequestService) RejectRequest(finReq models.FinRequest) error {
	r, err := s.repo.FinRequestById(finReq.Id)
	if err != nil {
		return err
	}
	if r.Status == models.RequestApproved || r.Status == models.RequestRejected {
		return errors.New("заявка уже обработана")
	}
	return s.repo.UpdateStatus(finReq.Id, models.RequestRejected)
}
