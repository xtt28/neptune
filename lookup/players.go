package lookup

import (
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/google/uuid"
	"github.com/xtt28/neptune/database/model"
	"gorm.io/gorm"
)

func GetOnlineOrOfflineUUID(db *gorm.DB, srv *server.Server, username string) (id uuid.UUID, online bool, err error) {
	p, ok := srv.PlayerByName(username)
	if ok {
		return p.UUID(), true, nil
	}

	id, err = GetOfflineUUID(db, username)
	return id, false, err
}

func GetOfflineUUID(db *gorm.DB, username string) (uuid.UUID, error) {
	var dest model.UserProfile
	err := db.First(&dest, &model.UserProfile{LastUsername: strings.ToLower(username)}).Error
	if err != nil {
		return uuid.Nil, err
	}
	return dest.UUID, nil
}
