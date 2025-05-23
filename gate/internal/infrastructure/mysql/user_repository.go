package mysql

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
	result := ur.db.Model(&domain.User{AccountId: id}).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (ur *UserRepository) InsertUser(user domain.User) error {
	result := ur.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
