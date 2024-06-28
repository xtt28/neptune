package permission

import (
	"log"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/permission/permlvl"
)

func SetPermission(id uuid.UUID, level permlvl.PermissionLevel, cache bool) {
	if cache {
		PermCache[id] = level
	}
	dest := &model.Permission{Subject: id}
	err := database.DB.FirstOrCreate(dest).Error

	if err != nil {
		log.Printf("could not get/create permission of %s: %s", id, err.Error())
	}

	dest.Level = level
	err = database.DB.Save(dest).Error
	if err != nil {
		log.Printf("could not save permission of %s: %s", id, err.Error())
	}
}
