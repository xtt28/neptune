package model

import "github.com/google/uuid"

type PvPStat struct {
	Subject uuid.UUID `gorm:"primaryKey"`
	Kills   uint
	Deaths  uint
}
