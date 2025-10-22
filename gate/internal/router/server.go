package router

import (
	"context"
	"fmt"
	"gate/docs"
	"gate/internal/config"
	"gate/internal/infrastructure/repository"
	"gate/internal/interface_adapter/handler"
	"gate/internal/interface_adapter/middleware"
	"gate/internal/usecase"
	"gate/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type GateServer struct {
	httpServer *http.Server
	tokenMaker token.Maker
	db         *gorm.DB
	rdb        *redis.Client
}

func NewServer(db *gorm.DB, rdb *redis.Client, config config.TokenConfig, port string) (*GateServer, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Secret, config.ExpireDur)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &GateServer{
		tokenMaker: tokenMaker,
		db:         db,
		rdb:        rdb,
	}

	server.setupRouter(port)
	return server, nil
}

func (s *GateServer) setupRouter(port string) {
	router := gin.New()
	router.Use(middleware.RequestID())
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // http://localhost:7690/swagger/index.html

	userUsecase := usecase.NewUserUsecase(repository.NewUserRepository(s.db), repository.NewAccountRepository(s.db))
	handler.RegisterUserRoutes(router, s.tokenMaker, userUsecase, true)

	accountUsecase := usecase.NewAccountUsecase(repository.NewAccountRepository(s.db))
	handler.RegisterAccountRoutes(router, s.tokenMaker, accountUsecase)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	s.httpServer = srv
}

func (s *GateServer) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

// Graceful shutdown support
func (s *GateServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
