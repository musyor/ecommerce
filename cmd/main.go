package main

import (
	"ecommerce/internal/config"
	"ecommerce/internal/db"
	"ecommerce/internal/handler"
	"ecommerce/internal/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// initial MySQL
	sql, err := db.NewMySQL(cfg)
	if err != nil {
		log.Fatalf("Failed to initial MySQL: %v ", err)
	}
	db.MySQLInstance = sql
	//initial Redis
	redis := db.NewRedis(cfg)

	db.RedisInstance = redis

	router := gin.Default()
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.LoginUser)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	s := server.NewServer(router)
	// 启动服务器
	log.Println("Server starting on :9090")
	if err := s.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
