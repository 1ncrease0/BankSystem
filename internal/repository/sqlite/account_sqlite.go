package sqlite

import (
	"FinanceSystem/internal/models"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountSQLite struct {
	db *sqlx.DB
}

func NewAccountSQLite(db *sqlx.DB) *AccountSQLite {
	return &AccountSQLite{db: db}
}

func (s *AccountSQLite) RollbackTransferMoney(fromAccountId, toAccountId int, amount float64) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var toBalance float64
	err = tx.Get(&toBalance, "SELECT balance FROM account WHERE id = ?", toAccountId)
	if err != nil {
		return err
	}
	if toBalance < amount {
		return errors.New("insufficient funds in recipient account to rollback transfer")
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = balance - ?, last_update = ? WHERE id = ?",
		amount,
		time.Now().UTC(),
		toAccountId,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = balance + ?, last_update = ? WHERE id = ?",
		amount,
		time.Now().UTC(),
		fromAccountId,
	)
	if err != nil {
		return err
	}

	return nil
}
func (s *AccountSQLite) TransferMoney(fromAccountId, toAccountId int, amount float64) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var fromStatus string
	err = tx.Get(&fromStatus, "SELECT status FROM account WHERE id = ?", fromAccountId)
	if err != nil {
		return err
	}
	if fromStatus != models.StatusAvailable {
		return errors.New("account is not available for transactions")
	}

	var fromBalance float64
	err = tx.Get(&fromBalance, "SELECT balance FROM account WHERE id = ?", fromAccountId)
	if err != nil {
		return err
	}
	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = balance - ?, last_update = ? WHERE id = ?",
		amount,
		time.Now().UTC(),
		fromAccountId,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = balance + ?, last_update = ? WHERE id = ?",
		amount,
		time.Now().UTC(),
		toAccountId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountSQLite) PlusMoney(accountId int, amount float64) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var currentBalance float64
	err = tx.Get(&currentBalance, "SELECT balance FROM account WHERE id = ?", accountId)
	if err != nil {
		return err
	}

	newBalance := currentBalance + amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = ?, last_update = ? WHERE id = ?",
		newBalance,
		time.Now().UTC(),
		accountId,
	)
	return err
}

func (s *AccountSQLite) MinusMoney(accountId int, amount float64) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var currentBalance float64
	err = tx.Get(&currentBalance, "SELECT balance FROM account WHERE id = ?", accountId)
	if err != nil {
		return err
	}

	newBalance := currentBalance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec(
		"UPDATE account SET balance = ?, last_update = ? WHERE id = ?",
		newBalance,
		time.Now().UTC(),
		accountId,
	)
	return err
}

func (s *AccountSQLite) CreateAccount(account models.Account) (int, error) {
	account.LastUpdate = time.Now().UTC()

	query := `INSERT INTO account (
			client_id,
			enterprise_id,
			bank_id,
			currency,
			balance,
			status,
			last_update
		) VALUES (
			:client_id,
			:enterprise_id,
			:bank_id,
			:currency,
			:balance,
			:status,
			:last_update
		)`
	result, err := s.db.NamedExec(query, account)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *AccountSQLite) DeleteAccount(id int) error {
	query := `DELETE FROM account WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountSQLite) Account(id int) (models.Account, error) {
	query := `SELECT * FROM account WHERE id = ?`
	var account models.Account
	err := s.db.Get(&account, query, id)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}

func (s *AccountSQLite) AccountsByClient(clientId int) ([]models.Account, error) {
	query := `
		SELECT *
		FROM account
		WHERE client_id = ?
	`
	var accounts []models.Account
	err := s.db.Select(&accounts, query, clientId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountSQLite) AccountsByEnterprise(enterpriseId int) ([]models.Account, error) {
	query := `
		SELECT *
		FROM account
		WHERE enterprise_id = ?
	`
	var accounts []models.Account
	err := s.db.Select(&accounts, query, enterpriseId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func isValidTransition(current, new string) bool {
	transitions := map[string][]string{
		models.StatusAvailable: {models.StatusFrozen, models.StatusBlocked},
		models.StatusFrozen:    {models.StatusAvailable, models.StatusBlocked},
		models.StatusBlocked:   {},
	}

	allowedTransitions, ok := transitions[current]
	if !ok {
		return false
	}

	for _, allowed := range allowedTransitions {
		if allowed == new {
			return true
		}
	}
	return false
}

func (s *AccountSQLite) UpdateStatus(accountId int, newStatus string) (err error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var currentStatus string
	err = tx.Get(&currentStatus, "SELECT status FROM account WHERE id = ?", accountId)
	if err != nil {
		return err
	}

	if !isValidTransition(currentStatus, newStatus) {
		return errors.New("invalid status transition")
	}

	_, err = tx.Exec(
		"UPDATE account SET status = ?  WHERE id = ?",
		newStatus,
		accountId,
	)
	return err
}
func (s *AccountSQLite) AccountsByBank(bankId int) ([]models.Account, error) {
	query := `
		SELECT *
		FROM account
		WHERE bank_id = ?
	`
	var accounts []models.Account
	err := s.db.Select(&accounts, query, bankId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
func (s *AccountSQLite) AccountsByClientAndBank(clientId, bankId int) ([]models.Account, error) {
	query := `
		SELECT *
		FROM account
		WHERE client_id = ? AND bank_id = ?
	`
	var accounts []models.Account
	err := s.db.Select(&accounts, query, clientId, bankId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountSQLite) AccountStatus(accountId int) (string, error) {
	var status string
	query := "SELECT status FROM account WHERE id = ?"
	err := s.db.Get(&status, query, accountId)
	if err != nil {
		return "", err
	}
	return status, nil
}

func (s *AccountSQLite) ClientBanks(clientId int) ([]models.Bank, error) {
	query := `
        SELECT DISTINCT b.id, b.name, b.bic 
        FROM account a
        JOIN bank b ON a.bank_id = b.id
        WHERE a.client_id = ?
    `

	var banks []models.Bank
	err := s.db.Select(&banks, query, clientId)
	if err != nil {
		return nil, err
	}

	return banks, nil
}
