package command

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/lookup"
	"github.com/xtt28/neptune/permission"
	"github.com/xtt28/neptune/permission/permlvl"
	"gorm.io/gorm"
)

type modHistoryCommandExec struct {
	srv     *server.Server
	Subject string `cmd:"subject"`
}

func (c modHistoryCommandExec) Run(source cmd.Source, output *cmd.Output) {
	p, ok := source.(*player.Player)
	if !ok {
		return
	}

	permLvl := permission.PermLevel(database.DB, p.UUID())
	if permLvl < permlvl.LvlModerator {
		permission.SendGateMessage(output, permlvl.LvlModerator)
		return
	}

	subject, _, err := lookup.GetOnlineOrOfflineUUID(database.DB, c.srv, c.Subject)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			output.Print(text.Colourf("<red>We couldn't find that player.</red>"))
		} else {
			output.Print(text.Colourf("<red>An error occurred while looking up that player.</red>"))
			log.Println(err.Error())
		}
		return
	}

	target := new([]*model.Punishment)
	res := database.DB.Order("created_at DESC").
		Where(&model.Punishment{Subject: subject}).
		Find(&target)

	if res.Error != nil {
		output.Print(text.Colourf("<red>An error occurred while looking up that history.</red>"))
		log.Println(res.Error.Error())
		return
	}

	bob := strings.Builder{}
	bob.WriteString(text.Colourf("\n<aqua>Enforcement history for <diamond>%s</diamond> <grey>(descending)</grey></aqua>\n", c.Subject))
	if len(*target) == 0 {
		bob.WriteString(text.Colourf("<green>This player has never been punished</green>\n"))
	}
	for i, v := range *target {
		bob.WriteString(text.Colourf("<diamond>%d. </diamond>", i + 1))
		bob.WriteString(text.Colourf("<aqua>%s</aqua> <dark-grey>-</dark-grey> ", v.Type))
		bob.WriteString(text.Colourf("<grey>%s</grey> (case #%d)\n", v.CreatedAt.Format(time.RFC822), v.ID))
		reason := v.Reason
		if reason == "" {
			reason = "<red>None provided</red>"
		}
		bob.WriteString(text.Colourf("<grey>Reason:</grey> <white>%s</white>\n", reason))
	}
	bob.WriteString("\n")
	p.Message(bob.String())
}
