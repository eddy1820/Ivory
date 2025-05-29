package usecase

import (
	"gate/internal/domain"
	"gate/internal/usecase/mocks"
	"gate/pkg/error_code"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestAccountUsecase_GetAccountInfoByAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountRepository(ctrl)

	expected := domain.Account{
		Account:           "eddy123",
		HashedPassword:    "secret",
		Email:             "eddy@gmail.com",
		CreatedAt:         time.Now(),
		PasswordChangedAt: time.Now(),
	}

	mockRepo.
		EXPECT().
		GetAccountInfoByAccount("eddy123").
		Return(expected, nil)

	uc := NewAccountUsecase(mockRepo)

	result, code := uc.GetAccountInfoByAccount("eddy123")

	assert.Equal(t, error_code.CodeOK, code)
	assert.Equal(t, expected.Account, result.Account)
}

func TestAccountUsecase_InsertAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountRepository(ctrl)
	uc := NewAccountUsecase(mockRepo)

	account := domain.Account{Account: "eddy123"}

	t.Run("account already exists", func(t *testing.T) {
		mockRepo.
			EXPECT().
			GetAccountInfoByAccount(account.Account).
			Return(domain.Account{Account: "eddy123"}, nil)

		code := uc.InsertAccount(account)
		assert.Equal(t, error_code.CodeAlreadyExists, code)
	})

	t.Run("error checking existence", func(t *testing.T) {
		mockRepo.
			EXPECT().
			GetAccountInfoByAccount(account.Account).
			Return(domain.Account{}, nil)

		code := uc.InsertAccount(account)

		assert.Equal(t, error_code.CodeAlreadyExists, code)
	})

	t.Run("account not found and insert success", func(t *testing.T) {
		mockRepo.
			EXPECT().
			GetAccountInfoByAccount(account.Account).
			Return(domain.Account{}, gorm.ErrRecordNotFound)

		mockRepo.
			EXPECT().
			InsertAccount(gomock.AssignableToTypeOf(domain.Account{})).
			Return(nil)

		code := uc.InsertAccount(account)

		assert.Equal(t, error_code.CodeOK, code)
	})
}
