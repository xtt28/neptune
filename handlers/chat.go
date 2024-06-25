package handlers

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/TwiN/go-away"
)

type ChatHandler struct {
	player.NopHandler
	p *player.Player
}

func newChatHandler(p *player.Player) *ChatHandler {
	return &ChatHandler{p: p}
}

func (m *ChatHandler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()

	*message = goaway.Censor(*message)

	fmt.Fprint(chat.Global, text.Colourf("<grey>%v <dark-grey>»</dark-grey> %v</grey>", m.p.Name(), *message))
}
