package lookup

import (
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/google/uuid"
	"github.com/xtt28/neptune/database/model"
	"gorm.io/gorm"
)

func GetOnlineOrOfflineUUID(db *gorm.DB, srv *server.Server, username string) (uuid.UUID, error) {
	p, ok := srv.PlayerByName(username)
	if ok {
		return p.UUID(), nil
	}

	return GetOfflineUUID(db, username)
}

func GetOfflineUUID(db *gorm.DB, username string) (uuid.UUID, error) {
	var dest model.UserProfile
	err := db.First(&dest, &model.UserProfile{LastUsername: strings.ToLower(username)}).Error
	if err != nil {
		return uuid.Nil, err
	}
	return dest.UUID, nil
}