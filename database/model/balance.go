package model

import "github.com/google/uuid"

type Balance struct {
	Subject uuid.UUID `gorm:"primaryKey"`
	Value uint64
}