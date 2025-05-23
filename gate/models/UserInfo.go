package models

import (
	"gate/internal/infrastructure/global"
)

func (this *UserInfo) TableName() string {
	return "user_info"
}

func (this *UserInfo) GetUserInfoById(id int64) (*UserInfo, error) {
	req := UserInfo{}
	result := global.DB.Model(&UserInfo{AccountId: id}).First(&req)
	if result.Error != nil {
		return nil, result.Error
	}

	return &req, nil
}

func (this *UserInfo) InsertUserInfo() error {
	result := global.DB.Create(this)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
