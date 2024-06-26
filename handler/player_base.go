package handler

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"gorm.io/gorm"
)

type BasePlayerHandler struct {
	player.NopHandler
	db *gorm.DB
	srv *server.Server
	p  *player.Player
}

func newBaseHandler(db *gorm.DB, srv *server.Server, p *player.Player) *BasePlayerHandler {
	return &BasePlayerHandler{db: db, srv: srv, p: p}
}
