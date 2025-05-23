package controller

import (
	"gate/internal/interface_adapter/middleware"
	"gate/internal/usecase"
	"gate/pkg/token"
	"gate/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type GetUserRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type SetUserRequest struct {
	Gender  string `form:"gender" json:"gender"`
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

type UserController struct {
	router      *gin.Engine
	userUsecase *usecase.UserUsecase
}

func NewUserController(router *gin.Engine, maker token.Maker, userUsecase *usecase.UserUsecase) *UserController {
	controller := &UserController{router: router, userUsecase: userUsecase}
	v1 := router.Group("/v1")
	userRouter := v1.Group("/user").Use(middleware.AuthMiddleware(maker))
	userRouter.POST("/", controller.SetUser)
	userRouter.GET("/:id", controller.GetUserById)
	return controller
}

// SetUser
// @Description 設定用戶資料
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Param 	Authorization 	header 		string 	true 	"token"
// @Param	gender			formData	string	true	"性別"
// @Param	name			formData	string	true	"姓名"
// @Param	address			formData	string	true	"地址"
// @Router /v1/user [post]
func (uc UserController) SetUser(ctx *gin.Context) {
	req := SetUserRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, router.ErrorResponse(err))
		return
	}
	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	err = uc.userUsecase.SetUser(payload.Username, req.Gender, req.Name, req.Address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, router.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

// GetById
// @Description 取得用戶
// @Tags 取得用戶
// @Success 200 string json{"code","method","path","id"}
// @Router /v1/user/{id} [get]
func (uc UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}
	user, err := uc.userUsecase.GetUserById(int64(idNum))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, router.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, &user)
}
