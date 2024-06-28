package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/form"
	"github.com/xtt28/neptune/permission/permlvl"
)

var kitCommand = cmd.New("kits", "Shows you a list of available kits.", []string{"kit"}, kitCommandExec{})

type kitCommandExec struct{}

func (c kitCommandExec) Run(source cmd.Source, output *cmd.Output) {
	if !RequireAtLeast(source, output, permlvl.LvlDefault) {
		return
	}
	p := source.(*player.Player)

	form.ShowKitSelector(p)
}
