package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     time.Duration
}

func NewPasetoMaker(symmetricKey string, duration time.Duration) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{paseto: paseto.NewV2(), symmetricKey: []byte(symmetricKey), duration: duration}

	return maker, nil
}

func (this PasetoMaker) CreateToken(username string) (string, error) {
	payload, err := NewPayload(username, this.duration)
	if err != nil {
		return "", err
	}
	return this.paseto.Encrypt(this.symmetricKey, payload, nil)
}

func (this PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := this.paseto.Decrypt(token, this.symmetricKey, payload, nil)
	if err != nil {
		return nil, errInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
