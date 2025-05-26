package controller

import (
	"gate/internal/domain"
	"gate/internal/infrastructure/global"
	"gate/internal/pkg/error_code"
	"gate/internal/pkg/token"
	"gate/internal/pkg/util"
	"gate/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountLoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

type AccountLoginResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

type AccountController struct {
	router         *gin.Engine
	tokenMaker     token.Maker
	accountUsecase *usecase.AccountUsecase
}

type CreateAccountReq struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type CreateAccountResponse struct {
	AccessToken string `json:"accessToken"`
}

func RegisterAccountRoutes(router *gin.Engine, maker token.Maker, accountUsecase *usecase.AccountUsecase) *AccountController {
	controller := &AccountController{router: router, tokenMaker: maker, accountUsecase: accountUsecase}
	v1 := router.Group("/v1")
	userRouter := v1.Group("/account")
	userRouter.POST("/login", controller.Login)
	userRouter.POST("/signIn", controller.SignIn)
	return controller
}

// SignIn
// @Description
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Param	email		formData	string	true	"信箱"
// @Router /v1/account/signIn [post]
func (ac AccountController) SignIn(ctx *gin.Context) {
	req := &CreateAccountReq{}

	err := ctx.Bind(&req)
	if err != nil {
		error_code.InvalidParams.SendResponse(ctx)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		error_code.InvalidParams.SendResponse(ctx)
		return
	}
	account := domain.Account{Account: req.Account, Email: req.Email, HashedPassword: hashedPassword}

	err = ac.accountUsecase.InsertAccount(account)
	if err != nil {
		error_code.InvalidParams.WithDetails(err.Error()).SendResponse(ctx)
		return
	}
	error_code.Success.SendResponse(ctx)
}

// Login
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Router /v1/account/login [post]
func (ac AccountController) Login(ctx *gin.Context) {
	req := AccountLoginRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		error_code.InvalidParams.SendResponse(ctx)
		return
	}

	account, err := ac.accountUsecase.GetAccountInfoByAccount(req.Account)
	if err != nil {
		error_code.ErrorUserNotExist.SendResponse(ctx)
		return
	}

	err = util.CheckPassword(req.Password, account.HashedPassword)
	if err != nil {
		error_code.InvalidParams.SendResponse(ctx)
		return
	}

	token, err := ac.tokenMaker.CreateToken(account.Account, global.TokenSetting.Expire)
	if err != nil {
		error_code.InvalidParams.SendResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, &AccountLoginResponse{AccessToken: token})
}
