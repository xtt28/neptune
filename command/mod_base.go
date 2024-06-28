package command

import (
	"database/sql"
	"log"
	"time"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/moderation"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
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

func modBaseRun(
	source cmd.Source,
	output *cmd.Output,
	pType moderation.PunishmentType,
	srv *server.Server,
	duration time.Duration,
	sub string,
	reason string,
) {
	if !RequireAtLeast(source, output, permlvl.LvlModerator) {
		return
	}
	p := source.(*player.Player)
	permLvl := permission.PermLevel(database.DB, p.UUID())

	subject, isOnline := lookup.GetOnlinePlayerCaseInsensitive(srv, sub)
	var id uuid.UUID
	if !isOnline {
		if pType == moderation.PunishmentTypeKick {
			output.Print(text.Colourf("<red>That player is not online.</red>"))
			return
		}
		offUUID, _, err := lookup.GetOnlineOrOfflineUUID(database.DB, srv, sub)
		if err != nil {
			output.Print(text.Colourf("<red>We couldn't find that player.</red>"))
			return
		}
		id = offUUID
	}

	if permission.PermLevel(database.DB, id) >= permLvl {
		output.Print(text.Colourf("<red>You may not issue a punishment against this player.</red>"))
		return
	}

	record := &model.Punishment{
		Issuer:  p.UUID(),
		Subject: id,
		Type:    pType,
		Reason:  reason,
	}
	if (pType == moderation.PunishmentTypeBan || pType == moderation.PunishmentTypeMute) && duration > 0 {
		record.ExpiresAt = sql.NullTime{Time: time.Now().Add(duration), Valid: true}
	}
	res := database.DB.Create(record)
	if res.Error != nil {
		output.Print(text.Colourf("<red>An error occurred while logging the punishment.</red>"))
		log.Println(res.Error.Error())
	}
	if isOnline {
		if pType == moderation.PunishmentTypeKick {
			subject.Disconnect(text.Colourf(messageFormats[moderation.PunishmentTypeKick], reason, record.ID))
		} else if pType == moderation.PunishmentTypeBan {
			var expiryStr string
			if record.ExpiresAt.Valid {
				expiryStr = record.ExpiresAt.Time.Format(time.RFC822)
			} else {
				expiryStr = "<red>PERMANENT</permanent>"
			}
			subject.Disconnect(text.Colourf(messageFormats[moderation.PunishmentTypeBan], expiryStr, reason, record.ID))
		}
	}

	output.Print(text.Colourf("<green>Punishment with case ID #%d successfully issued against %s</green>", record.ID, sub))
}
