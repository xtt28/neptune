package model

import "github.com/google/uuid"

type PvPStat struct {
	Subject uuid.UUID `gorm:"uniqueIndex"`
	Kills uint
	Deaths uint
}