package handler

import (
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

func (h *BasePlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	if src, ok := src.(entity.AttackDamageSource); ok {
		attacker, ok := src.Attacker.(*player.Player)
		if !ok {
			return
		}

		for _, player := range h.srv.Players() {
			player.Message(text.Colourf("<dark-grey>»</dark-grey> <emerald>%s</emerald> <grey>killed</grey> <redstone>%s</redstone> <dark-grey>(%.1f )</dark-grey>", attacker.Name(), h.p.Name(), attacker.Health() / 2))
		}
	}
}
