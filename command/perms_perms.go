package command

import (
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
)

type permsSetCommandExec struct{
	srv *server.Server
	Set cmd.SubCommand `cmd:"set"`
	Subject string `cmd:"subject"`
	Level string `cmd:"level"`
}

func (c permsSetCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	permLvl := permission.PermLevel(database.DB, p.UUID())
	if permLvl < permlvl.LvlAdmin {
		permission.SendGateMessage(output, permlvl.LvlAdmin)
		return
	}

	level, ok := permlvl.NameToLevel[strings.ToLower(c.Level)]
	if !ok {
		output.Print(text.Colourf("<red>This permission level is not recognized</red>"))
		return
	}
	player, ok := c.srv.PlayerByName(c.Subject)
	if ok {
		permission.SetPermission(player.UUID(), level, true)
	} else {
		parsedUUID, err := uuid.Parse(c.Subject)
		if err != nil {
			output.Print(text.Colourf("<red>For an offline player, specify a valid UUID</red>"))
			return
		}
		permission.SetPermission(parsedUUID, level, false)
	}
	output.Print(text.Colourf("<green>Done</green>"))
}
