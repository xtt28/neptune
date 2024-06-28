package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
)

func RequireAtLeast(source cmd.Source, output *cmd.Output, lvl permlvl.PermissionLevel) bool {
	p, ok := source.(*player.Player)
	if !ok {
		return false
	}

	permLvl := permission.PermLevel(database.DB, p.UUID())
	if permLvl < lvl {
		permission.SendGateMessage(output, lvl)
		return false
	}

	return true
}
