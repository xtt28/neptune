package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/menu"
)

var kitCommand = cmd.New("kits", "Shows you a list of available kits.", []string{"kit"}, kitCommandExec{})

type kitCommandExec struct{}

func (c kitCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	menu.ShowKitSelector(p)
}
