package db

import (
	"ecommerce/internal/config"
	"ecommerce/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB // 公开的字段
}

var MySQLInstance *MySQL

func NewMySQL(cfg *config.Config) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 迁移用户表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return &MySQL{DB: db}, nil
}
