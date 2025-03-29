package models

const (
	RoleOperator      = "operator"
	RoleManager       = "manager"
	RoleAdministrator = "administrator"
)

type BankEmployee struct {
	Id       int    `json:"-"                  db:"id"`
	UserName string `json:"user_name"          db:"username"`
	Role     string `json:"role"               db:"role"`
	BankId   int    `json:"bank_id"            db:"bank_id"`
	Password string `json:"password"           db:"password"`
}
