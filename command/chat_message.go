package command

import (
	goaway "github.com/TwiN/go-away"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/permission/permlvl"
)

type privateMessageCommandExec struct {
	srv *server.Server
	Target string `cmd:"target"`
	Message cmd.Varargs `cmd:"message"`
}

func (c privateMessageCommandExec) Run(source cmd.Source, output *cmd.Output) {
	if !RequireAtLeast(source, output, permlvl.LvlDefault) {
		return
	}
	p := source.(*player.Player)

	target, ok := lookup.GetOnlinePlayerCaseInsensitive(c.srv, c.Target)
	if !ok {
		output.Print(text.Colourf("<red>We couldn't find that player.</red>"))
		return
	}

	censored := goaway.Censor(string(c.Message))
	target.Message(text.Colourf("<grey>From %s: %s</grey>", p.Name(), censored))
	target.Message(text.Colourf("<grey>To %s: %s</grey>", c.Target, censored))
}
