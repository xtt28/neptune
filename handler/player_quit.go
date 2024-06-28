package handler

import "github.com/xtt28/neptune/game"

func (h *BasePlayerHandler) HandleQuit() {
	combat, ok := game.Combat.Combats[h.p];
	if !ok {
		return
	}

	game.RecordKill(combat.Attacker, h.p)
}