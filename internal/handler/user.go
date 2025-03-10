package handler

import (
	"ecommerce/internal/db"
	"ecommerce/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hash/crc32"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// 创建用户
	if err := db.MySQLInstance.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// 登录用户
func LoginUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码哈希处理
	hashedPassword := crc32.ChecksumIEEE([]byte(user.Password))
	user.Password = fmt.Sprintf("%d", hashedPassword)

	// 查询用户
	var storedUser model.User
	if err := db.MySQLInstance.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成会话ID
	sessionID := crc32.ChecksumIEEE([]byte(time.Now().String() + user.Username))

	// 将会话ID存储到Redis
	redisClient := db.RedisInstance.Client
	redisClient.Set(c, fmt.Sprintf("session:%d", sessionID), user.Username, time.Hour*24)

	c.SetCookie("session_id", fmt.Sprintf("%d", sessionID), 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    storedUser,
	})
}
