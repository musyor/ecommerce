package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex:idx_username;size:255;not null"`
	Email     string    `gorm:"uniqueIndex:idx_email;size:255;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
