package models

type Bank struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Bic  string `db:"bic"`
}

type FilteredBank struct {
	Bank   Bank
	Filter bool
}
