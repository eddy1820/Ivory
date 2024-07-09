package models

import (
	"fmt"
	"gate/global"
	"time"
)

type AccountInfo struct {
	Id                int64     `json:"id,omitempty"`
	Account           string    `json:"account,omitempty"`
	HashedPassword    string    `json:"hashedPassword,omitempty"`
	Email             string    `json:"email,omitempty"`
	PasswordChangedAt time.Time `json:"passwordChanged_at"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (this *AccountInfo) TableName() string {
	return "account_info"
}

func (this *AccountInfo) GetAccountInfoByAccount(account string) (*AccountInfo, error) {
	response := AccountInfo{}
	result := global.DB.Where("account = ?", account).First(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return &response, nil
}

func (this *AccountInfo) InsertAccount() error {
	info, _ := this.GetAccountInfoByAccount(this.Account)
	if info != nil {
		return fmt.Errorf("account is already exists")
	}
	this.CreatedAt = time.Now()
	this.PasswordChangedAt = time.Now()

	result := global.DB.Create(this)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
