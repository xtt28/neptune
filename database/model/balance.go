package model

import "github.com/google/uuid"

type Balance struct {
	Subject uuid.UUID `gorm:"uniqueIndex"`
	Value uint64
}