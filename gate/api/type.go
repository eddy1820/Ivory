package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	Method string `form:"method" json:"method"`
	Path   string `form:"path" json:"path"`
	Data   string `form:"data" json:"data"`
}

// GetById
// @Description 取得用戶
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @Param	id	path	int64	true	"用戶id"
// @Router /v1/type/{id} [get]
func (this *Server) GetById(c *gin.Context) {
	id := c.Param("id")
	//c.JSON(http.StatusOK, gin.H{
	//	"method": c.Request.Method,
	//	"path":   c.Request.URL.Path,
	//	"id":     id,
	//})

	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: id})
}

type ByParamsRequest struct {
	Id string `form:"id" binding:"required"`
}

// ByParams
// @Description 取得用戶
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @Param	id	query	int64	true	"用戶id"
// @Router /v1/type/byParams [get]
func (this *Server) ByParams(c *gin.Context) {
	req := ByParamsRequest{}

	err := c.ShouldBindQuery(&req)
	if err != nil {
		panic(err)
		return
	}

	id := c.Query("id")

	fmt.Println(id)

	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: req.Id})
}

// ByFormData
// @Description 取得用戶
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @Param	name	formData	string	true	"姓名"
// @Router /v1/type/byFormData [post]
func (this *Server) ByFormData(c *gin.Context) {
	name := c.PostForm("name")
	//c.JSON(http.StatusOK, gin.H{
	//	"method": c.Request.Method,
	//	"path":   c.Request.URL.Path,
	//	"name":   name,
	//})

	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: name})

}

// ByRaw
// @Description Raw data
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @param data body string true "json string"
// @Router /v1/type/byRaw [post]
func (this *Server) ByRaw(c *gin.Context) {
	data, _ := c.GetRawData()
	//c.JSON(http.StatusOK, gin.H{
	//	"method": c.Request.Method,
	//	"path":   c.Request.URL.Path,
	//	"data":   string(data),
	//})
	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: string(data)})
}

type ByJsonRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"  binding:"required,number"`
}

// ByJson
// @Description ByJson
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @param data body string true "json string"
// @Router /v1/type/byJson [post]
func (this *Server) ByJson(c *gin.Context) {
	req := ByJsonRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("err : " + err.Error())
		panic(err)
	}

	jsonString, _ := json.Marshal(req)
	fmt.Println("req : " + string(jsonString))
	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: string(jsonString)})
}

type ByUriRequest struct {
	Id   int64  `uri:"id" binding:"required,min=1"`
	Name string `uri:"name" binding:"required"`
}

// ByJson
// @Description ByJson
// @Tags API類型
// @Success 200 string json{"code","method","path","id"}
// @Param	id		path	int64	true	"用戶id"
// @Param	name	path	string	true	"用戶姓名"
// @Router /v1/type/{id}/{name} [get]
func (this *Server) ByUri(c *gin.Context) {
	req := ByUriRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Println("err : " + err.Error())
		panic(err)
	}

	c.JSON(http.StatusOK, &Response{Method: c.Request.Method, Path: c.Request.URL.Path, Data: strconv.FormatInt(req.Id, 10) + req.Name})
}
