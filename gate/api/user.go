package api

import (
	"gate/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type GetUserRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

// GetById
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	id	path	int64	true	"用戶id"
// @Router /v1/user/{id} [get]
func (this *Server) GetUserById(c *gin.Context) {
	req := GetUserRequest{}

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	model := models.UserInfo{}
	userInfo, err := model.GetUserInfoById(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &userInfo)
}
