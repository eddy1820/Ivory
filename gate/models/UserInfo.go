package models

import (
	"gate/global"
)

type UserInfo struct {
	Id        int64  `json:"id"`
	AccountId int64  `json:"accountId,omitempty" json:"account_id,omitempty"`
	Gender    string `json:"gender,omitempty" json:"gender,omitempty"`
	Name      string `json:"name,omitempty" json:"name,omitempty"`
	Address   string `json:"phone,omitempty" json:"address,omitempty"`
}

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

func (this *UserInfo) InsertUserInfo(info *UserInfo) error {
	result := global.DB.Create(info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
