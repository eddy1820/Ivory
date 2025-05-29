package handler

import (
	"bytes"
	"encoding/json"
	"gate/internal/domain"
	"gate/internal/interface_adapter/middleware"
	"gate/internal/usecase/mocks"
	"gate/internal/usecase/usecase_interface"
	"gate/pkg/error_code"
	"gate/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUserRouterAndHandler(t *testing.T, mockUsecase usecase_interface.UserUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	fakePayload := &token.Payload{Username: "eddy123"}
	maker, _ := token.NewPasetoMaker("12345678901234567890123456789012")

	r.Use(func(c *gin.Context) {
		c.Set(middleware.AuthorizationPayloadKey, fakePayload)
		c.Next()
	})

	RegisterUserRoutes(r, maker, mockUsecase, false)
	return r
}

func TestUserHandler_SetUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockUserUsecase(ctrl)
	mockUsecase.
		EXPECT().
		SetUser("eddy123", "male", "Eddy", "Taipei").
		Return(error_code.CodeOK)

	router := setupUserRouterAndHandler(t, mockUsecase)

	body := map[string]string{
		"gender":  "male",
		"name":    "Eddy",
		"address": "Taipei",
	}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/v1/user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token123")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUserHandler_GetUserById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockUserUsecase(ctrl)
	mockUsecase.
		EXPECT().
		GetUserById(int64(1)).
		Return(domain.User{Id: 1, Name: "Eddy"}, error_code.CodeOK)

	router := setupUserRouterAndHandler(t, mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/v1/user/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUserHandler_GetUserById_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockUserUsecase(ctrl)
	router := setupUserRouterAndHandler(t, mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/v1/user/abc", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestUserHandler_GetUserById_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockUserUsecase(ctrl)
	mockUsecase.
		EXPECT().
		GetUserById(int64(99)).
		Return(domain.User{}, error_code.CodeNotFound)

	router := setupUserRouterAndHandler(t, mockUsecase)

	req := httptest.NewRequest(http.MethodGet, "/v1/user/99", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
