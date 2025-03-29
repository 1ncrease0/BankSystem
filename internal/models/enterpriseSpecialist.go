package models

type EnterpriseSpecialist struct {
	Id           int    `db:"id"`
	UserName     string `db:"username"`
	EnterpriseId int    `db:"enterprise_id"`
	Password     string `db:"password"`
}
