package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

var pingCommand = cmd.New("ping", "Pong!", []string{}, pingCommandExec{})

type pingCommandExec struct{}

func (c pingCommandExec) Run(source cmd.Source, output *cmd.Output) {
	output.Print(text.Colourf("<green>Pong!</green>"))
}
