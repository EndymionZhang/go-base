package model

import "github.com/shopspring/decimal"

type Transaction struct {
	ID            int
	FromAccountId int
	FromAccount   Account
	ToAccountId   int
	ToAccount     Account
	Amount        decimal.Decimal `gorm:"type:decimal(20,8)"`
}
