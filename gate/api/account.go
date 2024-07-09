package api

import (
	"gate/global"
	"gate/models"
	"gate/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAccountReq struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type CreateAccountResponse struct {
	AccessToken string `json:"accessToken"`
}

// SignIn
// @Description
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Param	email		formData	string	true	"信箱"
// @Router /v1/account/signIn [post]
func (this *Server) SignIn(c *gin.Context) {
	req := &CreateAccountReq{}

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	model := models.AccountInfo{Account: req.Account, Email: req.Email, HashedPassword: hashedPassword}

	err = model.InsertAccount()

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}

type AccountLoginRequest struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

type AccountLoginResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

// Login
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	account		formData	string	true	"帳號"
// @Param	password	formData	string	true	"密碼"
// @Router /v1/account/login [post]
func (this *Server) Login(c *gin.Context) {
	req := AccountLoginRequest{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	model := models.AccountInfo{}

	account, err := model.GetAccountInfoByAccount(req.Account)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, account.HashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	token, err := this.tokenMaker.CreateToken(account.Account, global.TokenSetting.Expire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &AccountLoginResponse{AccessToken: token})
}
