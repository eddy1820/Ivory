package api

import (
	"fmt"
	"gate/docs"
	"gate/pkg/setting"
	"gate/pkg/token"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config *setting.TokenSettings) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Secret)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
	}

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("currency", validCurrency)
	//}

	server.setupRouter()
	return server, nil
}

func (this *Server) setupRouter() {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // http://localhost:7500/swagger/index.html

	{
		v1 := router.Group("/v1")
		typeRouter := v1.Group("/type")
		typeRouter.GET("/:id", this.GetById)
		typeRouter.GET("/byParams", this.ByParams)
		typeRouter.POST("/byFormData", this.ByFormData)
		typeRouter.POST("/byRaw", this.ByRaw)
		typeRouter.POST("/byJson", this.ByJson)
		typeRouter.GET("/:id/:name", this.ByUri)
	}
	{
		v1 := router.Group("/v1")
		redisRouter := v1.Group("/redis")
		redisRouter.GET("/string", this.String)
		redisRouter.GET("/hash", this.Hash)
		redisRouter.GET("/list", this.List)
		redisRouter.GET("/set", this.Set)
		redisRouter.GET("/sortedSet", this.SortedSet)
	}
	{
		v1 := router.Group("/v1")
		userRouter := v1.Group("/user")
		userRouter.POST("", this.Post)
		userRouter.GET("/:id", this.GetUserById)
	}

	//router.POST("/users", server.createUser)
	//router.POST("/users/login", server.loginUser)
	//router.POST("/tokens/renew_access", server.renewAccessToken)
	//
	//authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	//authRoutes.POST("/accounts", server.createAccount)
	//authRoutes.GET("/accounts/:id", server.getAccount)
	//authRoutes.GET("/accounts", server.listAccounts)
	//
	//authRoutes.POST("/transfers", server.createTransfer)

	this.router = router
}

// Start runs the HTTP server on a specific address.
func (this *Server) Start(address string) error {
	return this.router.Run(":" + address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
