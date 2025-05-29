package handler

import (
	"gate/internal/interface_adapter/middleware"
	"gate/internal/usecase/usecase_interface"
	error_code2 "gate/pkg/error_code"
	token2 "gate/pkg/token"
	"github.com/gin-gonic/gin"
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

type UserHandler struct {
	router      *gin.Engine
	userUsecase usecase_interface.UserUsecase
}

func RegisterUserRoutes(router *gin.Engine, maker token2.Maker, userUsecase usecase_interface.UserUsecase, enableAuth bool) *UserHandler {
	controller := &UserHandler{router: router, userUsecase: userUsecase}
	v1 := router.Group("/v1")
	userRouter := v1.Group("/user")
	if enableAuth {
		userRouter.Use(middleware.AuthMiddleware(maker))
	}
	userRouter.POST("", controller.SetUser)
	userRouter.GET("/:id", controller.GetUserById)
	return controller
}

// SetUser
// @Summary Set the user information
// @Description Set the user information
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param body body handler.SetUserRequest true "User info"
// @Success 200 {object} error_code.ErrorData
// @Router /v1/user [post]
func (uc UserHandler) SetUser(ctx *gin.Context) {
	req := SetUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token2.Payload)
	code := uc.userUsecase.SetUser(payload.Username, req.Gender, req.Name, req.Address)
	switch code {
	case error_code2.CodeNotFound:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	case error_code2.CodeDBError:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	error_code2.Success.SendResponse(ctx)
}

// GetUserById
// @Summary Get User information by user_id
// @Description Get User info
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Router /v1/user/{id} [get]
func (uc UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	user, code := uc.userUsecase.GetUserById(int64(idNum))
	switch code {
	case error_code2.CodeNotFound:
		error_code2.NotFound.SendResponse(ctx)
		return
	case error_code2.CodeDBError:
		error_code2.InvalidParams.SendResponse(ctx)
		return
	}
	error_code2.SuccessResponse(ctx, &user)
}
