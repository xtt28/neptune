package permission

import (
	"log"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/permission/permlvl"
)

func SetPermission(id uuid.UUID, level permlvl.PermissionLevel) {
	PermCache[id] = level
	err := database.DB.Where(&model.Permission{Subject: id}).Update("level", level).Error
	if err != nil {
		log.Printf("could not set permission of %s: %s", id, err.Error())
	}
}