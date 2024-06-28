package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountReq struct {
	Account  string `form:"account" json:"account" binding:"required,alphaunicode"`
	Password string `form:"password" json:"password" binding:"required,number"`
}

func (this *Server) Post(c *gin.Context) {
	req := &accountReq{}

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   req,
	})
}

// GetById
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param	id	path	int64	true	"用戶id"
// @Router /v1/user/:id [get]
func (this *Server) GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"id":     id,
	})
}
