package db

import (
	"database/sql"
	"ecommerce/internal/config"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(cfg *config.Config) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDatabase,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MySQL{db: db}, nil
}
