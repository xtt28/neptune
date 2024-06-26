package models

import (
	"github.com/google/uuid"
	"github.com/xtt28/neptune/permissions"
)

type Permission struct {
	Subject uuid.UUID `gorm:"uniqueIndex"`
	Level   permissions.PermissionLevel
}
