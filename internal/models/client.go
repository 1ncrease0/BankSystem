package models

const RoleClient = "client"

type Client struct {
	Id             int    `json:"-"                                     db:"id"`
	Name           string `json:"name"            binding:"required"    db:"name"`
	Surname        string `json:"surname"         binding:"required"    db:"surname"`
	Patronymic     string `json:"patronymic"      binding:"required"    db:"patronymic"`
	UserName       string `json:"user_name"       binding:"required"    db:"username"`
	PassportSeries string `json:"passport_series" binding:"required"    db:"passport_series"`
	PassportNumber string `json:"passport_number" binding:"required"    db:"passport_number"`
	IdNumber       string `json:"id_number"       binding:"required"    db:"id_number"`
	PhoneNumber    string `json:"phone_number"    binding:"required"    db:"phone_number"`
	Email          string `json:"email"           binding:"required"    db:"email"`
	Password       string `json:"password"        binding:"required"    db:"password"`
}
