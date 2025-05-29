package domain

import "time"

type Account struct {
	Id                int64     `json:"id"`
	Account           string    `json:"account"`
	HashedPassword    string    `json:"hashedPassword"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (this *Account) TableName() string {
	return "accounts"
}
