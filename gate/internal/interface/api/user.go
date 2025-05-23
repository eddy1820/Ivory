package api

import (
	"gate/internal/middleware"
	"gate/models"
	"gate/pkg/token"
	"gate/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetUser
// @Description 設定用戶資料
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param 	Authorization 	header 		string 	true 	"token"
// @Param	gender			formData	string	true	"性別"
// @Param	name			formData	string	true	"姓名"
// @Param	address			formData	string	true	"地址"
// @Router /v1/user [post]
func (this *router.Server) SetUser(c *gin.Context) {

	req := SetUserRequest{}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.errorResponse(err))
		return
	}

	payload := c.MustGet(middleware.authorizationPayloadKey).(*token.Payload)

	account := models.AccountInfo{}
	accountInfo, err := account.GetAccountInfoByAccount(payload.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.errorResponse(err))
		return
	}
	model := models.UserInfo{AccountId: accountInfo.Id, Gender: req.Gender, Name: req.Name, Address: req.Address}
	err = model.InsertUserInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &model)
}

// GetById
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Router /v1/user [get]
func (this *router.Server) GetUserById(c *gin.Context) {

	model := models.UserInfo{}
	userInfo, err := model.GetUserInfoById(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, router.errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &userInfo)
}
