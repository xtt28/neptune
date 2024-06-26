package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"gorm.io/gorm"
)

func PlayerHandler(db *gorm.DB) func(*player.Player) {
	return func(p *player.Player) {
		handleJoin(db, p)
		p.Handle(newBaseHandler(db, p))
	}
}
