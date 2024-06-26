package commands

import "github.com/df-mc/dragonfly/server/cmd"

func RegisterCommands() {
	cmd.Register(spawnCommand)
	cmd.Register(kitCommand)
}
