package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model `json:"gorm.Model"`
	Account    string `json:"account,omitempty"`
	Password   string `json:"password,omitempty"`
	Name       string `json:"name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
}

func (this *UserInfo) TableName() string {
	return "user_info"
}
