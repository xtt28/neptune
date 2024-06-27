package handler

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xtt28/neptune/game"
)

func (h *BasePlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	combat, ok := game.Combat.Combats[h.p]
	if !ok {
		game.RecordMiscDeath(h.p)
		return
	}

	game.RecordKill(combat.Attacker, h.p)
}
