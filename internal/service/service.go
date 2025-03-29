package service

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository"
	"log"
	"time"
)

type Authorization interface {
	CreateClient(client models.Client) (int, error)
	Client(id int) (models.Client, error)
	CreateBankEmployee(employee models.BankEmployee) (int, error)
	GenerateClientToken(username, password string) (string, error)
	GenerateEmployeeToken(username, password string) (string, error)
	//GenerateToken(username, password, role string) (string, error)
	EnterpriseSpecialist(id int) (models.EnterpriseSpecialist, error)
	ParseToken(accessToken string) (*TokenClaims, error)
	CreateEnterpriseSpecialist(es models.EnterpriseSpecialist) (int, error)
	GenerateEnterpriseSpecialistToken(username, password string) (string, error)
}

type Bank interface {
	Banks() ([]models.Bank, error)
	Bank(bankId int) (models.Bank, error)
	BanksByClient(clientId int) ([]models.Bank, error)
	FilteredByClient(clientId int) ([]models.FilteredBank, error)
}

type Account interface {
	CreateAccount(account models.Account) (int, error)
	DeleteAccount(accountId int) error
	Account(accountId int) (models.Account, error)
	AccountsByClient(clientId int) ([]models.Account, error)
	AccountsByBank(bankId int) ([]models.Account, error)
	AccountsByClientAndBank(clientId, bankId int) ([]models.Account, error)
	ChangeStatus(accountId int, newStatus string) error
	Status(accountId int) (string, error)
	TransferMoney(from, to int, amount float64) error
	PutMoney(accountId int, amount float64) error
	WithdrawMoney(accountId int, amount float64) error
	AccountsByEnterprise(enterpriseId int) ([]models.Account, error)
	CancelTransfer(logID, fromAccountId, toAccountId int, amount float64) error
}

type Deposit interface {
	CreateDeposit(dep models.Deposit) (int, error)
	GetDeposit(id int) (models.Deposit, error)
	GetDepositsByAccount(accountID int) ([]models.Deposit, error)
	DeleteDeposit(id int) error
	GetAllDeposits() ([]models.Deposit, error)
	UpdateAmount(depositID int, newAmount float64, lastUpdate time.Time) error
}

type Credit interface {
	CreateCredit(cr models.Credit) (int, error)
	GetCredit(id int) (models.Credit, error)
	GetCreditsByAccount(accountID int) ([]models.Credit, error)
	DeleteCredit(id int) error
	GetAllCredits() ([]models.Credit, error)
	UpdateRemaining(creditID int, newRemaining float64, updatedAt time.Time) error
}

type Installment interface {
	CreateInstallment(inst models.Installment) (int, error)
	GetInstallment(id int) (models.Installment, error)
	GetInstallmentsByAccount(accountID int) ([]models.Installment, error)
	DeleteInstallment(id int) error
	GetAllInstallments() ([]models.Installment, error)
	UpdateRemaining(installmentID int, newRemaining float64, updatedAt time.Time) error
}

type ClientFinance interface {
	ClientFinanceInfoByBank(clientId, bankId int) (models.ClientFinanceInfo, error)
}

type RegistrationRequest interface {
	CreateRegistrationRequest(reg models.RegistrationRequest) (int, error)
	RegistrationRequestsByBank(bankId int) ([]models.RegistrationRequest, error)
	UpdateStatus(registrationRequestId int, status string) error
	DeleteRegistrationRequest(registrationRequestId int) error

	ApproveRequest(reqId, bankId int) error
	RejectRequest(reqId, bankId int) error
}

type FinRequest interface {
	CreateFinRequest(req models.FinRequest) (int, error)
	GetFinRequestByID(id int) (models.FinRequest, error)
	GetFinRequestsByClient(clientID int) ([]models.FinRequest, error)
	GetFinRequestsByBank(bankID int) ([]models.FinRequest, error)
	ChangeStatus(finRequestID int, status string) error
	RemoveFinRequest(finRequestID int) error

	ApproveRequest(finReq models.FinRequest) error
	RejectRequest(finReq models.FinRequest) error
}
type SalaryProject interface {
	SalaryProject(id int) (models.SalaryProject, error)
	SalaryProjectsByBank(bankId int) ([]models.SalaryProject, error)
	SalaryProjectsByClient(clientId int) ([]models.SalaryProject, error)
	SalaryProjectByEnterprise(enterpriseId int) ([]models.SalaryProject, error)
	CreateSalaryProject(project models.SalaryProject) (int, error)
	DeleteSalaryProject(id int) error
	UpdateSalaryProject(project models.SalaryProject) error
}

type SalaryProjectRequest interface {
	CreateSalaryProjectRequest(req models.SalaryProjectRequest) (int, error)
	GetSalaryProjectRequest(id int) (models.SalaryProjectRequest, error)
	GetSalaryProjectRequestsByClient(clientId int) ([]models.SalaryProjectRequest, error)
	GetSalaryProjectRequestsByEnterprise(enterpriseId int) ([]models.SalaryProjectRequest, error)
	GetSalaryProjectRequestsByBank(bankId int) ([]models.SalaryProjectRequest, error)
	UpdateRequestStatus(id int, status string) error
	DeleteSalaryProjectRequest(id int) error

	ApproveRequest(requestID int, bankID int) error
	RejectRequest(requestID int, bankID int) error
}

type ActionLog interface {
	Create(log models.ActionLogUnit) (int, error)
	GetByID(id int) (models.ActionLogUnit, error)
	GetAll() ([]models.ActionLogUnit, error)
	GetByBankID(bankId int) ([]models.ActionLogUnit, error)
	Delete(id int) error
}

type Service struct {
	Authorization        Authorization
	Account              Account
	Bank                 Bank
	Deposit              Deposit
	Credit               Credit
	Installment          Installment
	ClientFinance        ClientFinance
	RegistrationRequest  RegistrationRequest
	FinRequest           FinRequest
	SalaryProject        SalaryProject
	SalaryProjectRequest SalaryProjectRequest
	ActionLog            ActionLog
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization:        NewAuthService(repository.Authorization),
		Account:              NewAccountService(repository.Account, repository.ActionLog),
		Bank:                 NewBankService(repository.Account, repository.Bank),
		Deposit:              NewDepositService(repository.Deposit),
		Installment:          NewInstallmentService(repository.Installment),
		Credit:               NewCreditService(repository.Credit),
		ClientFinance:        NewClientFinanceService(repository.Account, repository.Deposit, repository.Credit, repository.Installment),
		RegistrationRequest:  NewRegistrationRequestService(repository.RegistrationRequest, NewAccountService(repository.Account, repository.ActionLog)),
		FinRequest:           NewFinRequestService(repository.FinRequest, NewAccountService(repository.Account, repository.ActionLog), NewDepositService(repository.Deposit), NewCreditService(repository.Credit), NewInstallmentService(repository.Installment)),
		SalaryProject:        NewSalaryProjectService(repository.SalaryProject),
		SalaryProjectRequest: NewSalaryProjectRequestService(repository.SalaryProjectRequest, NewAccountService(repository.Account, repository.ActionLog), NewSalaryProjectService(repository.SalaryProject)),
		ActionLog:            NewActionLogService(repository.ActionLog),
	}
}

func (s *Service) StartPeriodicRecalculation(interval time.Duration, stop <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := s.RecalculateAccounts(); err != nil {
					log.Printf("Ошибка перерасчёта: %v", err)
				} else {
					log.Println("Перерасчёт успешно выполнен")
				}
			case <-stop:
				log.Println("Остановка периодического пересчёта")
				return
			}
		}
	}()
}

func (s *Service) RecalculateAccounts() error {
	now := time.Now()
	monthDuration := 30 * 24 * time.Hour

	deposits, err := s.Deposit.GetAllDeposits()
	if err != nil {
		return err
	}
	for _, dep := range deposits {
		if now.Sub(dep.UpdatedAt) >= monthDuration {
			monthlyInterestRate := dep.InterestRate / 100 / 12
			interest := dep.Amount * monthlyInterestRate
			newAmount := dep.Amount + interest

			if err := s.Deposit.UpdateAmount(dep.Id, newAmount, now); err != nil {
				log.Printf("Ошибка обновления депозита id=%d: %v", dep.Id, err)
			} else {
				log.Printf("Депозит id=%d обновлён: сумма изменена с %.2f на %.2f", dep.Id, dep.Amount, newAmount)
			}
		}
	}

	credits, err := s.Credit.GetAllCredits()
	if err != nil {
		return err
	}
	for _, cr := range credits {
		if cr.StartDate != nil && now.Sub(*cr.StartDate) >= monthDuration {
			monthlyInterestRate := cr.InterestRate / 100 / 12
			interest := cr.Remaining * monthlyInterestRate
			newRemaining := cr.Remaining + interest

			if err := s.Credit.UpdateRemaining(cr.Id, newRemaining, now); err != nil {
				log.Printf("Ошибка обновления кредита id=%d: %v", cr.Id, err)
			} else {
				log.Printf("Кредит id=%d обновлён: остаток изменён с %.2f на %.2f", cr.Id, cr.Remaining, newRemaining)
			}
		}
	}

	// Обработка рассрочек
	installments, err := s.Installment.GetAllInstallments()
	if err != nil {
		return err
	}
	for _, inst := range installments {
		if inst.StartDate != nil && now.Sub(*inst.StartDate) >= monthDuration {
			monthlyP := inst.Amount / float64(inst.TermMonths)
			newRemaining := inst.Remaining - monthlyP

			if err := s.Installment.UpdateRemaining(inst.Id, newRemaining, now); err != nil {
				log.Printf("Ошибка обновления рассрочки id=%d: %v", inst.Id, err)
			} else {
				log.Printf("Рассрочка id=%d обновлена: остаток изменён с %.2f на %.2f", inst.Id, inst.Remaining, newRemaining)
			}
		}
	}

	return nil
}
