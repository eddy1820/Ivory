package usecase

import (
	"errors"
	"gate/internal/domain"
	"gate/internal/usecase/mocks"
	"gate/pkg/error_code"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestUserUsecase_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockAccountRepo := mocks.NewMockAccountRepository(ctrl)

	expectedUser := domain.User{Id: 1, AccountId: 10, Name: "Eddy", Gender: "male", Address: "Taipei"}

	mockUserRepo.
		EXPECT().
		GetUserById(int64(1)).
		Return(expectedUser, nil)

	uc := NewUserUsecase(mockUserRepo, mockAccountRepo)

	user, code := uc.GetUserById(1)

	assert.Equal(t, error_code.CodeOK, code)
	assert.Equal(t, expectedUser, user)

}

func TestUserUsecase_SetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockAccountRepo := mocks.NewMockAccountRepository(ctrl)

	uc := NewUserUsecase(mockUserRepo, mockAccountRepo)

	username := "eddy123"
	account := domain.Account{Id: 10, Account: username}
	newUser := domain.User{AccountId: 10, Gender: "male", Name: "Eddy", Address: "Taipei"}

	t.Run("account not found", func(t *testing.T) {
		mockAccountRepo.
			EXPECT().
			GetAccountInfoByAccount(username).
			Return(domain.Account{}, gorm.ErrRecordNotFound)

		code := uc.SetUser(username, "male", "Eddy", "Taipei")
		assert.Equal(t, error_code.CodeNotFound, code)
	})

	t.Run("insert user failed", func(t *testing.T) {
		mockAccountRepo.
			EXPECT().
			GetAccountInfoByAccount(username).
			Return(account, nil)

		mockUserRepo.
			EXPECT().
			InsertUser(newUser).
			Return(errors.New("DB insert failed"))

		code := uc.SetUser(username, "male", "Eddy", "Taipei")
		assert.Equal(t, error_code.CodeDBError, code)
	})

	t.Run("successfully set user", func(t *testing.T) {
		mockAccountRepo.
			EXPECT().
			GetAccountInfoByAccount(username).
			Return(account, nil)

		mockUserRepo.
			EXPECT().
			InsertUser(newUser).
			Return(nil)

		code := uc.SetUser(username, "male", "Eddy", "Taipei")
		assert.Equal(t, error_code.CodeOK, code)
	})
}
