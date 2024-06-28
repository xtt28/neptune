package stats

import (
	"log"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
)

func GetStats(subject *player.Player) *model.PvPStat {
	subjectUUID := subject.UUID()

	dest := &model.PvPStat{
		Subject: subjectUUID,
	}
	err := database.DB.Where(dest).FirstOrCreate(dest).Error
	if err != nil {
		log.Printf("could not lookup balance of %s: %s\n", subjectUUID, err.Error())
		return dest
	}

	return dest
}
