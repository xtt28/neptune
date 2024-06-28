package model

import "github.com/google/uuid"

type UserProfile struct {
	UUID uuid.UUID `gorm:"primaryKey"`
	LastUsername string `gorm:"uniqueIndex"`
}