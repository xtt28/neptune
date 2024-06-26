package handler

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/game"
)

func (h *BasePlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	combat, ok := game.Combat.Combats[h.p]
	if !ok {
		fmt.Fprint(chat.Global, text.Colourf("<dark-grey>»</dark-grey> <emerald>%s</emerald> <grey>died</grey>", h.p.Name()))
		return
	}

	attacker := combat.Attacker

	fmt.Fprint(chat.Global, text.Colourf("<dark-grey>»</dark-grey> <emerald>%s</emerald> <grey>killed</grey> <redstone>%s</redstone> <dark-grey>(%.1f )</dark-grey>", attacker.Name(), h.p.Name(), attacker.Health() / 2))

	game.Combat.Clear(h.p)
}
