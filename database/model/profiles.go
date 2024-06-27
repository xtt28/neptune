package model

import "github.com/google/uuid"

type UserProfile struct {
	UUID uuid.UUID `gorm:"uniqueIndex"`
	LastUsername string `gorm:"uniqueIndex"`
}