package handlers

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/scoreboard"
	"gorm.io/gorm"
)

func handleJoin(db *gorm.DB, p *player.Player) {
	scoreboard.Render(p)

	p.SendTitle(title.New(text.Colourf("<diamond><bold>Welcome</bold></diamond>")).WithSubtitle(text.Colourf("<grey>to Neptune</grey>")))
}
