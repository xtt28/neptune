package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/game"
)

var spawnCommand = cmd.New("spawn", "Teleports you to the server's spawn point.", []string{}, spawnCommandExec{})

type spawnCommandExec struct{}

func (c spawnCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	game.SendToSpawn(p, false)
}
