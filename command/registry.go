package command

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

func RegisterCommands(srv *server.Server) {
	cmd.Register(spawnCommand)
	cmd.Register(kitCommand)
	cmd.Register(statsCommand)

	cmd.Register(cmd.New("permissions", "Manage user permissions on the server.", []string{"perms", "rank"}, permsSetCommandExec{srv: srv}))

	cmd.Register(cmd.New("message", "Send a private message to someone else.", []string{"pm", "msg", "tell", "t", "w", "whisper"}, privateMessageCommandExec{srv: srv}))
}
