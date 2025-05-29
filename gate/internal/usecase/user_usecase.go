package usecase

import (
	"errors"
	"gate/internal/domain"
	"gate/internal/usecase/repo_interface"
	"gate/pkg/error_code"
	"gorm.io/gorm"
)

type UserUsecase struct {
	userRepository    repo_interface.UserRepository
	accountRepository repo_interface.AccountRepository
}

func NewUserUsecase(userRepository repo_interface.UserRepository, accountRepository repo_interface.AccountRepository) *UserUsecase {
	return &UserUsecase{
		userRepository:    userRepository,
		accountRepository: accountRepository,
	}
}

func (uc *UserUsecase) GetUserById(id int64) (domain.User, int) {
	user, err := uc.userRepository.GetUserById(id)
	if err == nil {
		return user, error_code.CodeOK
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, error_code.CodeNotFound
	}
	return user, error_code.CodeDBError
}

func (uc *UserUsecase) SetUser(username string, gender string, name string, address string) int {
	account, err := uc.accountRepository.GetAccountInfoByAccount(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error_code.CodeNotFound
		}
		return error_code.CodeDBError
	}

	user := domain.User{AccountId: account.Id, Gender: gender, Name: name, Address: address}
	if err := uc.userRepository.InsertUser(user); err != nil {
		return error_code.CodeDBError
	}
	return error_code.CodeOK
}
