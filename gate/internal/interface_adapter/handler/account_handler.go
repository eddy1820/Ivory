package handler

import (
	"gate/internal/domain"
	"gate/internal/infrastructure/global"
	"gate/internal/usecase/usecase_interface"
	error_code2 "gate/pkg/error_code"
	"gate/pkg/token"
	"gate/pkg/util"
	"github.com/gin-gonic/gin"
)

type AccountLoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

type AccountLoginResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

type AccountHandler struct {
	router         *gin.Engine
	tokenMaker     token.Maker
	accountUsecase usecase_interface.AccountUsecase
}

type CreateAccountReq struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type CreateAccountResponse struct {
	AccessToken string `json:"accessToken"`
}

func RegisterAccountRoutes(router *gin.Engine, maker token.Maker, accountUsecase usecase_interface.AccountUsecase) *AccountHandler {
	controller := &AccountHandler{router: router, tokenMaker: maker, accountUsecase: accountUsecase}
	v1 := router.Group("/v1")
	userRouter := v1.Group("/account")
	userRouter.POST("/login", controller.Login)
	userRouter.POST("/register", controller.Register)
	return controller
}

// Register
// @Description Register
// @Tags Account
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Param	email		formData	string	true	"信箱"
// @Router /v1/account/register [post]
func (ac AccountHandler) Register(ctx *gin.Context) {
	req := &CreateAccountReq{}
	err := ctx.Bind(&req)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	account := domain.Account{Account: req.Account, Email: req.Email, HashedPassword: hashedPassword}
	code := ac.accountUsecase.InsertAccount(account)
	switch code {
	case error_code2.CodeAlreadyExists:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	case error_code2.CodeDBError:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	error_code2.Success.SendResponse(ctx)
}

// Login
// @Description Login
// @Tags Account
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Router /v1/account/login [post]
func (ac AccountHandler) Login(ctx *gin.Context) {
	req := AccountLoginRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}

	account, code := ac.accountUsecase.GetAccountInfoByAccount(req.Account)
	switch code {
	case error_code2.CodeNotFound:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	case error_code2.CodeDBError:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}

	err = util.CheckPassword(req.Password, account.HashedPassword)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}

	token, err := ac.tokenMaker.CreateToken(account.Account, global.TokenSetting.Expire)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	error_code2.SuccessResponse(ctx, &AccountLoginResponse{AccessToken: token})
}
