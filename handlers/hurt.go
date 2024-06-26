package handlers

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xtt28/neptune/config"
)

func (h *BasePlayerHandler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	if h.p.Position().Y() > config.NConfig.Locations.DmgThreshold {
		ctx.Cancel()
	}

	if _, ok := src.(entity.FallDamageSource); ok {
		ctx.Cancel()
	}

	if _, ok := src.(entity.VoidDamageSource); ok {
		*damage *= 20
	}
}
