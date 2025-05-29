package repo_interface

import "gate/internal/domain"

//go:generate mockgen -source=user_repository.go -destination=../../usecase/mocks/mock_user_repository.go -package=mocks
type UserRepository interface {
	GetUserById(id int64) (user domain.User, err error)
	InsertUser(user domain.User) error
}
