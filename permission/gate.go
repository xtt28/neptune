package permission

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/permission/permlvl"
)

func SendGateMessage(output *cmd.Output, minLvl permlvl.PermissionLevel) {
	output.Print(text.Colourf("<red><bold>NO PERMISSION</bold></red> <grey>You must be <red>%s</red> or higher to issue this command!</grey>", permlvl.LevelToName[minLvl]))
}
