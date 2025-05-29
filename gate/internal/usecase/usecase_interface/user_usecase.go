package usecase_interface

import "gate/internal/domain"

//go:generate mockgen -source=user_usecase.go -destination=../../usecase/mocks/mock_user_usecase.go -package=mocks
type UserUsecase interface {
	GetUserById(id int64) (domain.User, int)
	SetUser(username string, gender string, name string, address string) int
}
