package command

import (
	"log"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/moderation"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
)

const kickMessage = `<aqua><bold>NEPTUNE ENFORCEMENT</bold></aqua>
<grey>You were kicked from the server.
Reason: <white>%s</white>
Case ID: <white>#%d</white></grey>`

type modKickCommandExec struct {
	srv     *server.Server
	Subject string         `cmd:"subject"`
	Reason  string         `cmd:"reason"`
}

func (c modKickCommandExec) Run(source cmd.Source, output *cmd.Output) {
	if !RequireAtLeast(source, output, permlvl.LvlModerator) {
		return
	}
	p := source.(*player.Player)
	permLvl := permission.PermLevel(database.DB, p.UUID())

	subject, ok := lookup.GetOnlinePlayerCaseInsensitive(c.srv, c.Subject)
	if !ok {
		output.Print(text.Colourf("<red>We couldn't find that player.</red>"))
		return
	}
	if subject.UUID() == p.UUID() {
		output.Print(text.Colourf("<red>You can't kick yourself.</red>"))
		return
	}
	if permission.PermLevel(database.DB, subject.UUID()) >= permLvl {
		output.Print(text.Colourf("<red>You can't kick this player because their rank is higher than or equal to yours.</red>"))
		return
	}

	record := &model.Punishment{
		Issuer: p.UUID(),
		Subject: subject.UUID(),
		Type: moderation.PunishmentTypeKick,
		Reason: c.Reason,
	}
	res := database.DB.Create(record)
	if res.Error != nil {
		output.Print(text.Colourf("<red>An error occurred while logging the punishment. We will still kick the player.</red>"))
		log.Println(res.Error.Error())
	}
	subject.Disconnect(text.Colourf(kickMessage, c.Reason, record.ID))
	output.Print(text.Colourf("<green>Punishment #%d successfully issued against %s</green>", record.ID, c.Subject))
}
