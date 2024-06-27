package handler

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"gorm.io/gorm"
)

func PlayerHandler(db *gorm.DB, srv *server.Server) func(*player.Player) {
	return func(p *player.Player) {
		handleJoin(p)
		p.Handle(newBaseHandler(db, srv, p))
	}
}
