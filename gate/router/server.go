package router

import (
	"fmt"
	"gate/docs"
	"gate/internal/controller"
	"gate/internal/infrastructure/mysql"
	"gate/internal/infrastructure/setting"
	"gate/internal/usecase"
	"gate/pkg/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router     *gin.Engine
	tokenMaker token.Maker
	db         *gorm.DB
}

func NewServer(db *gorm.DB, config *setting.TokenSettings) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Secret)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
		db:         db,
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // http://localhost:7500/swagger/index.html

	userUsecase := usecase.NewUserUsecase(mysql.NewUserRepository(s.db), mysql.NewAccountRepository(s.db))
	controller.NewUserController(router, s.tokenMaker, userUsecase)

	//{
	//	v1 := router.Group("/v1")
	//	userRouter := v1.Group("/account")
	//	userRouter.POST("/login", s.Login)
	//	userRouter.POST("/signIn", s.SignIn)
	//}

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(":" + address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
