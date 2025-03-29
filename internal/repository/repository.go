package repository

import (
	"FinanceSystem/internal/models"
	"FinanceSystem/internal/repository/sqlite"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
// clientsTable = "client"
// bankTable    = "bank"
// accountTable      = "account"
// bankEmployeeTable = "bankEmployee"
// registrationReqTable = "registration_request"
// finRequestTable="fin_request"
// enterpriseSpecialistTable = "enterprise_specialist"
)

type Authorization interface {
	CreateClient(client models.Client) (int, error)
	Client(username, password string) (models.Client, error)
	ClientById(clientId int) (models.Client, error)

	CreateBankEmployee(employee models.BankEmployee) (int, error)
	BankEmployee(username, password string) (models.BankEmployee, error)

	CreateEnterpriseSpecialist(es models.EnterpriseSpecialist) (int, error)
	EnterpriseSpecialist(username, password string) (models.EnterpriseSpecialist, error)
	EnterpriseSpecialistById(id int) (models.EnterpriseSpecialist, error)
}

type Bank interface {
	Bank(name string) (models.Bank, error)
	AllBanks() ([]models.Bank, error)
	BankById(bankId int) (models.Bank, error)
}

type Account interface {
	CreateAccount(account models.Account) (int, error)
	DeleteAccount(id int) error
	Account(id int) (models.Account, error)
	AccountsByClient(clientId int) ([]models.Account, error)
	AccountsByEnterprise(enterpriseId int) ([]models.Account, error)
	AccountsByBank(bankId int) ([]models.Account, error)
	AccountsByClientAndBank(clientId, bankId int) ([]models.Account, error)

	ClientBanks(clientId int) ([]models.Bank, error)

	AccountStatus(accountId int) (string, error)
	UpdateStatus(accountId int, status string) error
	PlusMoney(accountId int, amount float64) error
	MinusMoney(accountId int, amount float64) error
	RollbackTransferMoney(fromAccountId, toAccountId int, amount float64) error
	TransferMoney(fromAccountId, toAccountId int, amount float64) error
}

type Deposit interface {
	CreateDeposit(dep models.Deposit) (int, error)
	Deposit(id int) (models.Deposit, error)
	DepositsByAccount(accountId int) ([]models.Deposit, error)
	DeleteDeposit(id int) error
	GetAllDeposits() ([]models.Deposit, error)
	UpdateAmount(depositID int, newAmount float64, lastUpdate time.Time) error
}

type Credit interface {
	CreateCredit(cr models.Credit) (int, error)
	Credit(id int) (models.Credit, error)
	CreditsByAccount(accountId int) ([]models.Credit, error)
	DeleteCredit(id int) error
	GetAllCredits() ([]models.Credit, error)
	UpdateRemaining(creditID int, newRemaining float64, updatedAt time.Time) error
}
type Installment interface {
	CreateInstallment(inst models.Installment) (int, error)
	Installment(id int) (models.Installment, error)
	InstallmentsByAccount(accountId int) ([]models.Installment, error)
	DeleteInstallment(id int) error
	GetAllInstallments() ([]models.Installment, error)
	UpdateRemaining(installmentID int, newRemaining float64, updatedAt time.Time) error
}

type RegistrationRequest interface {
	CreateRegistrationRequest(reg models.RegistrationRequest) (int, error)
	RegistrationRequestsByBank(bankId int) ([]models.RegistrationRequest, error)
	UpdateStatus(registrationRequestId int, status string) error
	DeleteRegistrationRequest(registrationRequestId int) error
	RegistrationRequestById(id int) (models.RegistrationRequest, error)
}

type FinRequest interface {
	CreateFinRequest(req models.FinRequest) (int, error)
	FinRequestById(id int) (models.FinRequest, error)
	FinRequestsByClient(clientId int) ([]models.FinRequest, error)
	FinRequestsByBank(bankId int) ([]models.FinRequest, error)
	UpdateStatus(finRequestId int, status string) error
	DeleteFinRequest(finRequestId int) error
}

type Enterprise interface {
	CreateEnterprise(enterprise models.Enterprise) (int, error)
	EnterpriseById(id int) (models.Enterprise, error)
	EnterpriseByUNP(unp string) (models.Enterprise, error)
}

type EnterpriseSpecialist interface {
	CreateEnterpriseSpecialist(es models.EnterpriseSpecialist) (int, error)
	EnterpriseSpecialist(username, password string) (models.EnterpriseSpecialist, error)
	EnterpriseSpecialistById(id int) (models.EnterpriseSpecialist, error)
}
type SalaryProject interface {
	CreateSalaryProject(salaryProject models.SalaryProject) (int, error)
	GetSalaryProjectById(id int) (models.SalaryProject, error)
	GetSalaryProjectsByEnterprise(enterpriseAccountId int) ([]models.SalaryProject, error)
	GetSalaryProjectsByClient(clientAccountId int) ([]models.SalaryProject, error)
	DeleteSalaryProject(id int) error
}

type SalaryProjectRequest interface {
	CreateSalaryProjectRequest(req models.SalaryProjectRequest) (int, error)
	GetSalaryProjectRequestById(id int) (models.SalaryProjectRequest, error)
	GetSalaryProjectRequestsByClient(clientAccountId int) ([]models.SalaryProjectRequest, error)
	GetSalaryProjectRequestsByEnterprise(enterpriseAccountId int) ([]models.SalaryProjectRequest, error)
	UpdateSalaryProjectRequestStatus(id int, status string) error
	DeleteSalaryProjectRequest(id int) error
	GetSalaryProjectRequestsByBank(bankId int) ([]models.SalaryProjectRequest, error)
}
type Repository struct {
	Authorization        Authorization
	Bank                 Bank
	Account              Account
	Deposit              Deposit
	Credit               Credit
	Installment          Installment
	RegistrationRequest  RegistrationRequest
	FinRequest           FinRequest
	Enterprise           Enterprise
	EnterpriseSpecialist EnterpriseSpecialist
	SalaryProject        SalaryProject
	SalaryProjectRequest SalaryProjectRequest
	ActionLog            ActionLog
}

type ActionLog interface {
	Create(log models.ActionLogUnit) (int, error)
	GetByID(id int) (models.ActionLogUnit, error)
	GetAll() ([]models.ActionLogUnit, error)
	GetByBankID(bankId int) ([]models.ActionLogUnit, error)
	Delete(id int) error
	GetByConsumer(consumerId int) ([]models.ActionLogUnit, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:        sqlite.NewAuthSQLite(db),
		Bank:                 sqlite.NewBankSQLite(db),
		Account:              sqlite.NewAccountSQLite(db),
		Deposit:              sqlite.NewDepositSQLite(db),
		Credit:               sqlite.NewCreditSQLite(db),
		Installment:          sqlite.NewInstallmentSQLite(db),
		RegistrationRequest:  sqlite.NewRegistrationRequestSQLite(db),
		FinRequest:           sqlite.NewFinRequestSQLite(db),
		Enterprise:           sqlite.NewEnterpriseSQLite(db),
		EnterpriseSpecialist: sqlite.NewEnterpriseSpecialistSQLite(db),
		SalaryProject:        sqlite.NewSalaryProjectSQLite(db),
		SalaryProjectRequest: sqlite.NewSalaryProjectRequestSQLite(db),
		ActionLog:            sqlite.NewActionLogSQLite(db),
	}
}
