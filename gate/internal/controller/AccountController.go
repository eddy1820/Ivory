package controller

import (
	"gate/internal/domain"
	"gate/internal/infrastructure/global"
	"gate/internal/usecase"
	"gate/pkg/token"
	"gate/pkg/util"
	"gate/router"
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

func NewAccountController(router *gin.Engine, maker token.Maker, accountUsecase *usecase.AccountUsecase) *AccountController {
	controller := &AccountController{router: router, accountUsecase: accountUsecase}
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
func (ac AccountController) SignIn(c *gin.Context) {
	req := &CreateAccountReq{}

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, router.ErrorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.ErrorResponse(err))
		return
	}
	account := domain.Account{Account: req.Account, Email: req.Email, HashedPassword: hashedPassword}

	err = ac.accountUsecase.InsertAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Login
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Router /v1/account/login [post]
func (ac AccountController) Login(c *gin.Context) {
	req := AccountLoginRequest{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, router.errorResponse(err))
		return
	}

	model := models.AccountInfo{}

	account, err := model.GetAccountInfoByAccount(req.Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, router.errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, account.HashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, router.errorResponse(err))
		return
	}

	token, err := this.tokenMaker.CreateToken(account.Account, global.TokenSetting.Expire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &AccountLoginResponse{AccessToken: token})
}
