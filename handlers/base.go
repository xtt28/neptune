package handlers

import (
	"github.com/df-mc/dragonfly/server/player"
	"gorm.io/gorm"
)

type BasePlayerHandler struct {
	player.NopHandler
	db *gorm.DB
	p  *player.Player
}

func newBaseHandler(db *gorm.DB, p *player.Player) *BasePlayerHandler {
	return &BasePlayerHandler{db: db, p: p}
}