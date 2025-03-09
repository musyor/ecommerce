package main

import (
	"context"
	"ecommerce/internal/config"
	"ecommerce/internal/db"
	"ecommerce/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
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
	//initial Redis
	redis := db.NewRedis(cfg)

	//初始化kafka生产者
	producter := db.NewProducer(cfg, "test")

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	router.GET("/mysql", func(c *gin.Context) {
		rows, err := sql.DB.Query("SELECT * FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"message": "MySQL connection successfully",
		})
	})
	router.GET("/redis", func(c *gin.Context) {
		pong, err := redis.Client.Ping(context.Background()).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": pong,
		})
	})
	router.GET("/kafka", func(c *gin.Context) {
		message := kafka.Message{
			Value: []byte("Test message"),
		}

		err := producter.Writer.WriteMessages(context.Background(), message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Kafka message sent successfully",
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})
	s := server.NewServer(router)
	// 启动服务器
	log.Println("Server starting on :9090")
	if err := s.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
