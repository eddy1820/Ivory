package domain

import "time"

type Account struct {
	Id                int64     `json:"id,omitempty"`
	Account           string    `json:"account,omitempty"`
	HashedPassword    string    `json:"hashedPassword,omitempty"`
	Email             string    `json:"email,omitempty"`
	PasswordChangedAt time.Time `json:"passwordChanged_at"`
	CreatedAt         time.Time `json:"createdAt"`
}
