package usecase

import (
	"fmt"
	"gate/internal/domain"
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
	if err != nil {
		return fmt.Errorf("account is already exists")
	}
	account.CreatedAt = time.Now()
	account.PasswordChangedAt = time.Now()

	err = ac.accountRepository.InsertAccount(account)
	if err != nil {
		return err
	}
	return nil
}
