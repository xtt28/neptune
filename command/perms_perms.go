package command

import (
	"errors"
	"log"
	"strings"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
	"gorm.io/gorm"
)

type permsSetCommandExec struct {
	srv     *server.Server
	Set     cmd.SubCommand `cmd:"set"`
	Subject string         `cmd:"subject"`
	Level   string         `cmd:"level"`
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
	id, _, err := lookup.GetOnlineOrOfflineUUID(database.DB, c.srv, strings.ToLower(c.Subject))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			output.Print(text.Colourf("<red>This player has never played on Neptune.</red>"))
		} else {
			output.Print(text.Colourf("<red>An internal error occurred and we couldn't find this player. More info in console.</red>"))
			log.Println(err.Error())
		}
		return
	}
	permission.SetPermission(id, level, true)
	output.Print(text.Colourf("<green>Done!</green>"))
}
