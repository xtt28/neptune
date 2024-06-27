package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/form"
)

var statsCommand = cmd.New("stats", "Shows you your game statistics.", []string{}, statsCommandExec{})

type statsCommandExec struct{}

func (c statsCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	form.ShowStats(p)
}
