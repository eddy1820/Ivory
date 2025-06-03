package usecase

import (
	"errors"
	"gate/internal/domain"
	"gate/internal/error_code"
	"gate/internal/usecase/repo_interface"
	"gorm.io/gorm"
	"time"
)

type AccountUsecase struct {
	accountRepository repo_interface.AccountRepository
}

func NewAccountUsecase(accountRepository repo_interface.AccountRepository) *AccountUsecase {
	return &AccountUsecase{
		accountRepository: accountRepository,
	}
}

func (ac *AccountUsecase) GetAccountInfoByAccount(account string) (domain.Account, int) {
	acc, err := ac.accountRepository.GetAccountInfoByAccount(account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Account{}, error_code.CodeNotFound
		}
		return domain.Account{}, error_code.CodeDBError
	}
	return acc, error_code.CodeOK
}

func (ac *AccountUsecase) InsertAccount(account domain.Account) int {
	_, err := ac.accountRepository.GetAccountInfoByAccount(account.Account)
	if err == nil {
		return error_code.CodeAlreadyExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return error_code.CodeDBError
	}
	account.CreatedAt = time.Now()
	account.PasswordChangedAt = time.Now()
	if err := ac.accountRepository.InsertAccount(account); err != nil {
		return error_code.CodeDBError
	}
	return error_code.CodeOK
}
