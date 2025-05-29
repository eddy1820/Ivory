package usecase_interface

import "gate/internal/domain"

//go:generate mockgen -source=account_usecase.go -destination=../../usecase/mocks/mock_account_usecase.go -package=mocks
type AccountUsecase interface {
	GetAccountInfoByAccount(account string) (domain.Account, int)
	InsertAccount(account domain.Account) int
}
