package repository

import (
	"gate/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserById(id int64) (user domain.User, err error) {
	result := ur.db.Where("account_id = ?", id).First(&user)
	return user, result.Error
}

func (ur *UserRepository) InsertUser(user domain.User) error {
	result := ur.db.Create(&user)
	return result.Error
}
