package handlers

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/database/models"
	"gorm.io/gorm"
)

func handleJoin(db *gorm.DB, p *player.Player) {
	perm := models.Permission{}
	db.First(&perm, models.Permission{Subject: p.UUID()})
	p.Messagef("your permission level is %d", perm.Level)
}
