package economy

import (
	"log"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/scoreboard"
)

func AddBits(target *player.Player, amount uint64) {
	current := GetBitsBalance(target)
	new := current + amount

	err := database.DB.Model(&model.Balance{}).Where(&model.Balance{Subject: target.UUID()}).Update("value", new).Error
	if err == nil {
		BitsCache[target.UUID()] = new
	} else {
		log.Println(err)
	}
	
	if amount > 0 {
		target.Message(text.Colourf("<grey><aqua>+%d<aqua> bits</grey>", amount))
	}
	scoreboard.Render(target, new)
}

func SubtractBits(target *player.Player, amount uint64) {
	AddBits(target, -amount)
}