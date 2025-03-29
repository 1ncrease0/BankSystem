package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"sort"
)

type BankService struct {
	accountRepo repository.Account
	bankRepo    repository.Bank
}

func NewBankService(accountRepo repository.Account, bankRepo repository.Bank) *BankService {
	return &BankService{
		accountRepo: accountRepo,
		bankRepo:    bankRepo}
}

func (s *BankService) BanksByClient(clientId int) ([]models.Bank, error) {
	return s.accountRepo.ClientBanks(clientId)
}

func (s *BankService) Bank(bankId int) (models.Bank, error) {
	bank, err := s.bankRepo.BankById(bankId)
	if err != nil {
		return models.Bank{}, err
	}
	return bank, nil
}

func (s *BankService) Banks() ([]models.Bank, error) {
	banks, err := s.bankRepo.AllBanks()
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (s *BankService) FilteredByClient(clientId int) ([]models.FilteredBank, error) {
	allBanks, err := s.bankRepo.AllBanks()
	if err != nil {
		return nil, err
	}

	clientBanks, err := s.accountRepo.ClientBanks(clientId)
	if err != nil {
		return nil, err
	}

	clientBanksMap := make(map[int]bool)
	for _, b := range clientBanks {
		clientBanksMap[b.Id] = true
	}

	var result []models.FilteredBank
	for _, bank := range allBanks {
		fb := models.FilteredBank{
			Bank:   bank,
			Filter: clientBanksMap[bank.Id],
		}
		result = append(result, fb)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Filter && !result[j].Filter
	})
	return result, nil
}
