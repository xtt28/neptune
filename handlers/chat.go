package handlers

import (
	"fmt"

	goaway "github.com/TwiN/go-away"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database/models"
	"github.com/xtt28/neptune/permissions"
)

func (m *BasePlayerHandler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()

	*message = goaway.Censor(*message)

	format := permissions.LevelToChatFormat[models.PermLevel(m.db, m.p.UUID())]
	fmt.Fprint(chat.Global, text.Colourf(format, m.p.Name(), *message))
}
