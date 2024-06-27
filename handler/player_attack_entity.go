package handler

import (
	"math/rand"

	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xtt28/neptune/game"
	"github.com/xtt28/neptune/game/kit"
)

func (m *BasePlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	p, ok := e.(*player.Player);
	if !ok {
		return
	}

	game.Combat.RecordHit(p, m.p)

	weapon, _ := m.p.HeldItems()
	if val, _ := weapon.Value(kit.ItemAbilityKey); val == kit.JupiterAbilityValue {
		if rand.Intn(4) == 0 { // 20% chance
			lightning := entity.NewLightningWithDamage(e.Position(), 6, false, 0)
			m.p.ShowEntity(lightning)
		}
	}
}
