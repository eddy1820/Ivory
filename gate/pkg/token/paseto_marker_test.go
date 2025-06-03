package token

import (
	"gate/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	duration := time.Minute
	maker, err := NewPasetoMaker(util.RandomString(32), duration)
	require.NoError(t, err)
	username := util.RandomOwner()

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotZero(t, payload.Id)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {

	duration := time.Minute

	maker, err := NewPasetoMaker(util.RandomString(32), -duration)
	require.NoError(t, err)
	username := util.RandomOwner()
	token, err := maker.CreateToken(username)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errExpiredToken.Error())
	require.Nil(t, payload)
}
