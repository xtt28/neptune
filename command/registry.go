package command

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

func RegisterCommands(srv *server.Server) {
	cmd.Register(spawnCommand)
	cmd.Register(kitCommand)
	
	cmd.Register(cmd.New("permissions", "Manage user permissions on the server.", []string{"perms", "rank"}, permsSetCommandExec{srv: srv}))
}
