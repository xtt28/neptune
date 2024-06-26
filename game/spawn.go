package game

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/config"
)

func SendToSpawn(p *player.Player, sendTitle bool) {
	spawn := config.NConfig.Locations.Spawn
	p.Teleport(mgl64.Vec3{spawn[0], spawn[1], spawn[2]})
	p.SetGameMode(world.GameModeAdventure)

	if sendTitle {
		p.SendTitle(title.New(text.Colourf("<red><bold>You died</bold></red>")).WithSubtitle(text.Colourf("<grey>You have respawned.</grey>")))
	}
}
