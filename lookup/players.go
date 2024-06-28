package lookup

import (
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"github.com/xtt28/neptune/database/model"
	"gorm.io/gorm"
)

// Taken from sliceutil in package github.com/df-mc/dragonfly/server/internal/sliceutil
func searchValue[A any, S ~[]A](v S, f func(a A) bool) (a A, ok bool) {
	for _, val := range v {
		if f(val) {
			return val, true
		}
	}
	return
}

func GetOnlinePlayerCaseInsensitive(srv *server.Server, username string) (*player.Player, bool) {
	return searchValue(srv.Players(), func(p *player.Player) bool {
		return strings.EqualFold(p.Name(), username)
	})
}

func GetOnlineOrOfflineUUID(db *gorm.DB, srv *server.Server, username string) (id uuid.UUID, online bool, err error) {
	p, ok := GetOnlinePlayerCaseInsensitive(srv, username)
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

func OfflineUUIDToUsername(db *gorm.DB, id uuid.UUID) (string, error) {
	var dest model.UserProfile
	err := db.First(&dest, &model.UserProfile{UUID: id}).Error
	if err != nil {
		return "", err
	}
	return dest.LastUsername, nil
}
