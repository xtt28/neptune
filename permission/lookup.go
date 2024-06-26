package permission

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/permission/permlvl"
	"gorm.io/gorm"
)

var PermCache = map[uuid.UUID]permlvl.PermissionLevel{}

func PermLevel(db *gorm.DB, playerID uuid.UUID) permlvl.PermissionLevel {
	cacheLvl, ok := PermCache[playerID]
	if ok {
		return cacheLvl
	}

	perm := model.Permission{}
	err := db.First(&perm, model.Permission{Subject: playerID}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(&model.Permission{Subject: playerID, Level: permlvl.LvlDefault}).Error
			if err != nil {
				log.Printf("could not create permissions for %s: %s", playerID, err.Error())
			}
			PermCache[playerID] = permlvl.LvlDefault
		} else {
			log.Printf("could not get user permissions for %s: %s", playerID, err.Error())
		}

		return permlvl.LvlDefault
	}

	PermCache[playerID] = perm.Level
	return perm.Level
}
