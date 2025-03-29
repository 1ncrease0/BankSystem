package sqlite

import (
	"FinanceSystem/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const bankTable = "bank"

type BankSQLite struct {
	db *sqlx.DB
}

func NewBankSQLite(db *sqlx.DB) *BankSQLite {
	return &BankSQLite{db: db}
}

func (r *BankSQLite) BankById(bankId int) (models.Bank, error) {
	var bank models.Bank
	query := fmt.Sprintf("SELECT id, name, bic FROM %s WHERE id=?", bankTable)
	err := r.db.Get(&bank, query, bankId)
	return bank, err
}

func (r *BankSQLite) Bank(name string) (models.Bank, error) {
	var bank models.Bank
	query := fmt.Sprintf("SELECT id, name, bic FROM %s WHERE name=?", bankTable)
	err := r.db.Get(&bank, query, name)
	return bank, err
}

func (r *BankSQLite) AllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	query := fmt.Sprintf("SELECT id, name, bic FROM %s", bankTable)
	err := r.db.Select(&banks, query)
	return banks, err
}
