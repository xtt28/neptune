package handlers

import "github.com/df-mc/dragonfly/server/player"

func PlayerHandler() func(*player.Player) {
	return func(p *player.Player) {
		p.Handle(newChatHandler(p))
	}
}