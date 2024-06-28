package punishments

import (
	"log"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/moderation"
)

var messageFormats = map[moderation.PunishmentType]string{
	"kick": `<aqua><bold>NEPTUNE ENFORCEMENT</bold></aqua>
<grey>You were kicked from the server.
Reason: <white>%s</white>
Case ID: <white>#%d</white></grey>`,
	"ban": `<aqua><bold>NEPTUNE ENFORCEMENT</bold></aqua>
<grey>You are banned from Neptune until <white>%s</white>.
Reason: <white>%s</white>
Case ID: <white>#%d</white></grey>`,
	"mute": `<aqua><bold>NEPTUNE ENFORCEMENT</bold></aqua>
<grey>You are muted from server chat until <white>%s</white>.
Reason: <white>%s</white>
Case ID: <white>#%d</white></grey>`,
}

func GenerateMessage(punishment model.Punishment) string {
	format := messageFormats[punishment.Type]
	switch punishment.Type {
	case moderation.PunishmentTypeKick:
		return text.Colourf(format, punishment.Reason, punishment.ID)
	case moderation.PunishmentTypeBan, moderation.PunishmentTypeMute:
		var expiryStr string
		if punishment.ExpiresAt.Valid {
			expiryStr = punishment.ExpiresAt.Time.Format(time.RFC822)
		} else {
			expiryStr = "<red>PERMANENT</permanent>"
		}
		return text.Colourf(format, expiryStr, punishment.Reason, punishment.ID)
	}
	return ""
}

func GetActive(p *player.Player, pType moderation.PunishmentType) *[]*model.Punishment {
	dest := new([]*model.Punishment)
	err := database.DB.
		Find(
			&dest,
			"type = ? AND subject = ? AND (expires_at >= ? OR expires_at IS NULL)",
			pType, p.UUID().String(), time.Now(),
		).Error
	if err != nil {
		log.Println(err)
	}
	return dest
}
