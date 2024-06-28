package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/moderation"
)

type Punishment struct {
	ID        uint `gorm:"primaryKey"`
	Issuer    uuid.UUID
	Subject   uuid.UUID
	Type      moderation.PunishmentType
	Reason    string
	CreatedAt time.Time
	ExpiresAt sql.NullTime
}
