package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
)

func (h *BasePlayerHandler) HandleItemDrop(ctx *event.Context, e world.Entity) {
	ctx.Cancel()
}