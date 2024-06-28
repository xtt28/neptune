package handler

import (
	"github.com/df-mc/dragonfly/server/event"
)

func (h *BasePlayerHandler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	ctx.Cancel()
}