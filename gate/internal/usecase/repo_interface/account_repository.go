package repo_interface

import "gate/internal/domain"

//go:generate mockgen -source=account_repository.go -destination=../../usecase/mocks/mock_account_repository.go -package=mocks
type AccountRepository interface {
	GetAccountInfoByAccount(account string) (domain.Account, error)
	InsertAccount(account domain.Account) error
}
