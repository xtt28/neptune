package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/moderation"
)

type Punishment struct {
	Issuer uuid.UUID
	Subject uuid.UUID
	Type moderation.PunishmentType
	CreatedAt time.Time
	ExpiresAt sql.NullTime
}