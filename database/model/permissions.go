package model

import (
	"github.com/google/uuid"
	"github.com/xtt28/neptune/permission"
)

type Permission struct {
	Subject uuid.UUID `gorm:"uniqueIndex"`
	Level   permission.PermissionLevel
}
