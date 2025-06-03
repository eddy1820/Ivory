package router

import (
	"fmt"
	"gate/docs"
	"gate/internal/config"
	"gate/internal/infrastructure/repository"
	"gate/internal/interface_adapter/handler"
	"gate/internal/usecase"
	"gate/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router     *gin.Engine
	tokenMaker token.Maker
	db         *gorm.DB
	rdb        *redis.Client
}

func NewServer(db *gorm.DB, rdb *redis.Client, config config.TokenConfig) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Secret, config.ExpireDur)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
		db:         db,
		rdb:        rdb,
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // http://localhost:7500/swagger/index.html

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(s.db), repository.NewAccountRepository(s.db))
	handler.RegisterUserRoutes(router, s.tokenMaker, userUsecase, true)

	accountUsecase := usecase.NewAccountUsecase(repository.NewAccountRepository(s.db))
	handler.RegisterAccountRoutes(router, s.tokenMaker, accountUsecase)

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(":" + address)
}
