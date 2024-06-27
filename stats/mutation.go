package stats

import (
	"log"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/economy/econlookup"
	"github.com/xtt28/neptune/scoreboard"
)

func Increment(target *player.Player, col string) {
	stats := GetStats(target)
	var stat uint
	if col == "kills" {
		stats.Kills++
		stat = stats.Kills
	} else if col == "deaths" {
		stats.Deaths++
		stat = stats.Deaths
	}

	err := database.DB.Model(&model.PvPStat{}).Where(&model.PvPStat{Subject: target.UUID()}).Update(col, stat).Error
	if err != nil {
		log.Println(err.Error())
	}

	scoreboard.Render(target, stats, econlookup.GetBitsBalance(target))
}

func AddKill(target *player.Player) {
	Increment(target, "kills")
}

func AddDeath(target *player.Player) {
	Increment(target, "deaths")
}
