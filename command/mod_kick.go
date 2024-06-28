package command

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/xtt28/neptune/moderation"
)

type modKickCommandExec struct {
	srv     *server.Server
	Subject string         `cmd:"subject"`
	Reason  string         `cmd:"reason"`
}

func (c modKickCommandExec) Run(source cmd.Source, output *cmd.Output) {
	modBaseRun(source, output, moderation.PunishmentTypeKick, c.srv, 0, c.Subject, c.Reason)
}
