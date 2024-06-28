package command

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
)

const kickMessage = `<aqua><bold>NEPTUNE ENFORCEMENT</bold></aqua>
<grey>You were kicked from the server.
Reason: <white>%s</white>

You may rejoin the server now to continue playing.</grey>`

type modKickCommandExec struct {
	srv     *server.Server
	Subject string         `cmd:"subject"`
	Reason  string         `cmd:"reason"`
}

func (c modKickCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	permLvl := permission.PermLevel(database.DB, p.UUID())
	if permLvl < permlvl.LvlModerator {
		permission.SendGateMessage(output, permlvl.LvlModerator)
		return
	}

	subject, ok := lookup.GetOnlinePlayerCaseInsensitive(c.srv, c.Subject)
	if !ok {
		output.Print(text.Colourf("<red>We couldn't find that player.</red>"))
		return
	}
	if subject.UUID() == p.UUID() {
		output.Print(text.Colourf("<red>You can't kick yourself.</red>"))
		return
	}
	if permission.PermLevel(database.DB, subject.UUID()) >= permLvl {
		output.Print(text.Colourf("<red>You can't kick this player because their rank is higher than or equal to yours.</red>"))
		return
	}

	subject.Disconnect(text.Colourf(kickMessage, c.Reason))
	output.Print(text.Colourf("<green>Punishment successfully issued against %s</green>", c.Subject))
}
