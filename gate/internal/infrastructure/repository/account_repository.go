package repository

import (
	"gate/internal/domain"
	"gorm.io/gorm"
	"time"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (ar *AccountRepository) GetAccountInfoByAccount(account string) (domain.Account, error) {
	data := domain.Account{}
	result := ar.db.Where("account = ?", account).First(&data)
	return data, result.Error
}

func (ar *AccountRepository) InsertAccount(account domain.Account) error {
	account.CreatedAt = time.Now()
	account.PasswordChangedAt = time.Now()
	result := ar.db.Create(&account)
	return result.Error
}
