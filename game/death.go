package game

import (
	"fmt"
	"math/rand"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/economy"
)

func RecordKill(attacker, target *player.Player) {
	attacker.SendPopup(text.Colourf("<grey>You killed <redstone>%s</redstone></grey>", target.Name()))
	fmt.Fprint(chat.Global, text.Colourf("<dark-grey>»</dark-grey> <emerald>%s</emerald> <grey>killed</grey> <redstone>%s</redstone> <dark-grey>(%.1f )</dark-grey>", attacker.Name(), target.Name(), attacker.Health() / 2))
	economy.AddBits(attacker, uint64(rand.Intn(6) + 5))

	Combat.Clear(target)
}

func RecordMiscDeath(victim *player.Player) {
	fmt.Fprint(chat.Global, text.Colourf("<dark-grey>»</dark-grey> <emerald>%s</emerald> <grey>died</grey>", victim.Name()))
}