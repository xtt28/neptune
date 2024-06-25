package handlers

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/config"
	"github.com/xtt28/neptune/scoreboard"
	"gorm.io/gorm"
)

func handleJoin(db *gorm.DB, p *player.Player) {
	scoreboard.Render(p)

	spawn := config.NConfig.Locations.Spawn
	p.Teleport(mgl64.Vec3{spawn[0], spawn[1], spawn[2]})
	p.SetGameMode(world.GameModeAdventure)

	p.SendTitle(title.New(text.Colourf("<diamond><bold>Welcome</bold></diamond>")).WithSubtitle(text.Colourf("<grey>to Neptune</grey>")))
}
