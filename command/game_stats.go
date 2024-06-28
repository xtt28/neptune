package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/form"
	"github.com/xtt28/neptune/permission/permlvl"
)

var statsCommand = cmd.New("stats", "Shows you your game statistics.", []string{}, statsCommandExec{})

type statsCommandExec struct{}

func (c statsCommandExec) Run(source cmd.Source, output *cmd.Output) {
	if !RequireAtLeast(source, output, permlvl.LvlDefault) {
		return
	}
	p := source.(*player.Player)
	form.ShowStats(p)
}
