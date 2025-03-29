package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
)

type ClientFinanceService struct {
	accountRepo     repository.Account
	depositRepo     repository.Deposit
	creditRepo      repository.Credit
	installmentRepo repository.Installment
}

func NewClientFinanceService(accRepo repository.Account, depRepo repository.Deposit, credRepo repository.Credit, instRepo repository.Installment) *ClientFinanceService {
	return &ClientFinanceService{
		accountRepo:     accRepo,
		depositRepo:     depRepo,
		creditRepo:      credRepo,
		installmentRepo: instRepo,
	}
}

func (s *ClientFinanceService) ClientFinanceInfoByBank(clientId, bankId int) (models.ClientFinanceInfo, error) {
	accounts, err := s.accountRepo.AccountsByClientAndBank(clientId, bankId)
	if err != nil {
		return models.ClientFinanceInfo{}, err
	}

	var deposits []models.Deposit
	var credits []models.Credit
	var installments []models.Installment

	for _, acc := range accounts {
		d, _ := s.depositRepo.DepositsByAccount(acc.Id)
		deposits = append(deposits, d...)

		c, _ := s.creditRepo.CreditsByAccount(acc.Id)
		credits = append(credits, c...)

		i, _ := s.installmentRepo.InstallmentsByAccount(acc.Id)
		installments = append(installments, i...)
	}

	return models.ClientFinanceInfo{
		Accounts:     accounts,
		Deposits:     deposits,
		Credits:      credits,
		Installments: installments,
	}, nil
}
