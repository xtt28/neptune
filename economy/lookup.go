package economy

import (
	"log"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
)

var BitsCache = map[uuid.UUID]uint64{}

func GetBitsBalance(subject *player.Player) uint64 {
	subjectUUID := subject.UUID()
	if cachedBal, ok := BitsCache[subjectUUID]; ok {
		return cachedBal
	}
	
	dest := &model.Balance{
		Subject: subjectUUID,
	}
	err := database.DB.Where(dest).FirstOrCreate(dest).Error
	if err != nil {
		log.Printf("could not lookup balance of %s: %s\n", subjectUUID, err.Error())
		return 0
	}

	BitsCache[subjectUUID] = dest.Value
	return dest.Value
}