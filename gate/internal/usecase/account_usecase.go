package usecase

import (
	"errors"
	"fmt"
	"gate/internal/domain"
	"gorm.io/gorm"
	"time"
)

type AccountUsecase struct {
	accountRepository AccountRepository
}

func NewAccountUsecase(accountRepository AccountRepository) *AccountUsecase {
	return &AccountUsecase{
		accountRepository: accountRepository,
	}
}

func (ac *AccountUsecase) GetAccountInfoByAccount(account string) (res domain.Account, err error) {
	res, err = ac.accountRepository.GetAccountInfoByAccount(account)
	if err != nil {
		return
	}
	return
}

func (ac *AccountUsecase) InsertAccount(account domain.Account) error {
	_, err := ac.accountRepository.GetAccountInfoByAccount(account.Account)
	if err == nil {
		return fmt.Errorf("account %s already exists", account.Account)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check account existence: %w", err)
	}
	account.CreatedAt = time.Now()
	account.PasswordChangedAt = time.Now()
	return ac.accountRepository.InsertAccount(account)
}
