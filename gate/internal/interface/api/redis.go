package api

import (
	"context"
	"fmt"
	"gate/internal/infrastructure/global"
	"gate/router"

	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// String
// @Description String
// @Tags Redis
// @Success 200 string json{"code","method","path","id"}
// @Param	data	query	string	true	"data"
// @Router /v1/redis/string [get]
func (this *router.Server) String(c *gin.Context) {
	data := c.Query("data")
	ctx := context.Background()
	err := global.Redis.Set(ctx, "KEY", data, 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	value, err := global.Redis.Get(ctx, "KEY").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("get KEY : " + value)

	value2, err := global.Redis.GetSet(ctx, "KEY", "CCC").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("get : " + value2 + " and set : " + "CCC")

	value3, err := global.Redis.Do(ctx, "get", "KEY").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("do get KEY : " + value3.(string))

	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   data,
	})
}

// Hash
// @Description Hash
// @Tags Redis
// @Success 200 string json{"code","method","path","id"}
// @Param	data	query	string	true	"data"
// @Router /v1/redis/hash [get]
func (this *router.Server) Hash(c *gin.Context) {
	data := c.Query("data")

	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   data,
	})
}

// List
// @Description List
// @Tags Redis
// @Success 200 string json{"code","method","path","id"}
// @Param	data	query	string	true	"data"
// @Router /v1/redis/list [get]
func (this *router.Server) List(c *gin.Context) {
	data := c.Query("data")
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   data,
	})
}

// Set
// @Description Set
// @Tags Redis
// @Success 200 string json{"code","method","path","id"}
// @Param	data	query	string	true	"data"
// @Router /v1/redis/set [get]
func (this *router.Server) Set(c *gin.Context) {
	data := c.Query("data")
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   data,
	})
}

// SortedSet
// @Description SortedSet
// @Tags Redis
// @Success 200 string json{"code","method","path","id"}
// @Param	data	query	string	true	"data"
// @Router /v1/redis/sortedSet [get]
func (this *router.Server) SortedSet(c *gin.Context) {
	data := c.Query("data")
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"data":   data,
	})
}
