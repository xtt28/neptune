package command

import (
	"time"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/moderation"
)

type modBanCommandExec struct {
	srv      *server.Server
	Subject  string `cmd:"subject"`
	Reason   string `cmd:"reason"`
	Duration string `cmd:"duration" optional:""`
}

func (c modBanCommandExec) Run(source cmd.Source, output *cmd.Output) {
	var duration time.Duration
	if c.Duration != "" {
		d, err := time.ParseDuration(c.Duration)
		if err != nil {
			output.Error(text.Colourf("<red>Please specify a valid duration (e.g. 7d)</red>"))
			return
		}
		duration = d
	}
	modBaseRun(source, output, moderation.PunishmentTypeBan, c.srv, duration, c.Subject, c.Reason)
}
