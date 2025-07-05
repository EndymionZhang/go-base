package model

import "github.com/shopspring/decimal"

type Account struct {
	ID      int
	Name    string          `gorm:"size:50"`
	Balance decimal.Decimal `gorm:"type:decimal(20,8)"`
}
