package handler

import (
	"bytes"
	"encoding/json"
	"gate/internal/domain"
	"gate/internal/usecase/mocks"
	"gate/internal/usecase/usecase_interface"
	"gate/pkg/response"
	"gate/pkg/token"
	"gate/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupAccountRouter(usecase usecase_interface.AccountUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	maker, _ := token.NewPasetoMaker("12345678901234567890123456789012", 10*time.Minute)
	RegisterAccountRoutes(router, maker, usecase)
	return router
}

func TestAccountHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockAccountUsecase(ctrl)
	router := setupAccountRouter(mockUsecase)

	body := CreateAccountReq{
		Account:  "eddy",
		Password: "strongpwd",
		Email:    "eddy@example.com",
	}

	mockUsecase.EXPECT().
		InsertAccount(gomock.Any()).
		Return(0)

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/v1/account/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestAccountHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockAccountUsecase(ctrl)

	router := setupAccountRouter(mockUsecase)

	hashedPwd, _ := util.HashPassword("strongpwd")

	mockUsecase.EXPECT().
		GetAccountInfoByAccount("eddy").
		Return(domain.Account{
			Account:        "eddy",
			HashedPassword: hashedPwd,
			Email:          "eddy@example.com",
		}, 0)

	body := AccountLoginRequest{
		Account:  "eddy",
		Password: "strongpwd",
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/v1/account/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	response := response.APIResponse[AccountLoginResponse]{}
	_ = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NotEmpty(t, response.Data.AccessToken)
}
