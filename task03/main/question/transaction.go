package question

import (
	"errors"
	"fmt"
	"github.com/endymion/go-base/task03/main/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account = model.Account
type Transaction = model.Transaction

func RunTrade(db *gorm.DB) {
	//createAccounts(db)
	trade(db)
}

func createAccounts(db *gorm.DB) {
	err := db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		fmt.Println("AutoMigrate error", err)
	}
	account := &Account{Name: "张三", Balance: decimal.NewFromInt(1000)}
	db.Create(account)
	account = &Account{Name: "李四", Balance: decimal.NewFromInt(1000)}
	db.Create(account)
	account = &Account{Name: "王五", Balance: decimal.NewFromInt(1000)}
	db.Create(account)
}

func trade(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount Account
		first := db.Debug().Where("name = ?", "张三").Where("id = ?", 1).First(&fromAccount)
		if first.RowsAffected < 1 || first.Error != nil {
			return errors.New("张三账户数据错误")
		}
		var toAccount Account
		db.Debug().Where("name = ?", "王五").Where("id = ?", 3).First(&toAccount)
		if first.RowsAffected < 1 || first.Error != nil {
			return errors.New("王五账户数据错误")
		}
		if fromAccount.Balance.Cmp(decimal.NewFromInt(100)) < 0 {
			return errors.New("余额不足")
		}
		fromAccount.Balance.Sub(decimal.NewFromInt(100))
		toAccount.Balance.Add(decimal.NewFromInt(100))
		db.Debug().Save(&fromAccount)
		db.Debug().Save(&toAccount)
		transaction := &Transaction{FromAccount: fromAccount, ToAccount: toAccount, Amount: decimal.NewFromInt(100)}
		db.Debug().Create(transaction)
		return nil
	})
	if err != nil {
		return
	}
}
