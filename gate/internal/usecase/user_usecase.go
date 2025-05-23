package usecase

import (
	"gate/internal/domain"
)

type UserRepository interface {
	GetUserById(id int64) (user domain.User, err error)
	InsertUser(user domain.User) error
}

type AccountRepository interface {
	GetAccountInfoByAccount(account string) (domain.Account, error)
	InsertAccount(account domain.Account) error
}

type UserUsecase struct {
	userRepository    UserRepository
	accountRepository AccountRepository
}

func NewUserUsecase(userRepository UserRepository, accountRepository AccountRepository) *UserUsecase {
	return &UserUsecase{
		userRepository:    userRepository,
		accountRepository: accountRepository,
	}
}

func (uc *UserUsecase) GetUserById(id int64) (domain.User, error) {
	return uc.userRepository.GetUserById(id)
}

func (uc *UserUsecase) SetUser(username string, gender string, name string, address string) error {
	account := domain.Account{}
	account, err := uc.accountRepository.GetAccountInfoByAccount(username)
	if err != nil {
		return err
	}
	user := domain.User{AccountId: account.Id, Gender: gender, Name: name, Address: address}
	err = uc.userRepository.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
