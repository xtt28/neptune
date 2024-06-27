package handler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/economy"
	"github.com/xtt28/neptune/game"
	"github.com/xtt28/neptune/scoreboard"
)

func handleJoin(p *player.Player) {
	p.EnableInstantRespawn()
	scoreboard.Render(p, economy.GetBitsBalance(p))
	
	game.SendToSpawn(p, false)
	
	p.SendTitle(title.New(text.Colourf("<diamond><bold>Welcome</bold></diamond>")).WithSubtitle(text.Colourf("<grey>to Neptune</grey>")))
}
