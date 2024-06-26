package model

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/xtt28/neptune/permission"
	"gorm.io/gorm"
)

var PermCache = map[uuid.UUID]permission.PermissionLevel{}

func PermLevel(db *gorm.DB, playerID uuid.UUID) permission.PermissionLevel {
	cacheLvl, ok := PermCache[playerID]
	if ok {
		return cacheLvl
	}

	perm := Permission{}
	err := db.First(&perm, Permission{Subject: playerID}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(&Permission{Subject: playerID, Level: permission.LvlDefault}).Error
			if err != nil {
				log.Printf("could not create permissions for %s: %s", playerID, err.Error())
			}
			PermCache[playerID] = permission.LvlDefault
		} else {
			log.Printf("could not get user permissions for %s: %s", playerID, err.Error())
		}

		return permission.LvlDefault
	}

	PermCache[playerID] = perm.Level
	return perm.Level
}
