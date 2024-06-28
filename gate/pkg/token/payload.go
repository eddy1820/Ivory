package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	errExpiredToken = errors.New("token has expired")
	errInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{Id: tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return &payload, nil
}

func (this Payload) Valid() error {
	if time.Now().After(this.ExpiredAt) {
		return errExpiredToken
	}
	return nil
}
