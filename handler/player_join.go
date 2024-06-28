package handler

import (
	"log"
	"strings"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database"
	"github.com/xtt28/neptune/database/model"
	"github.com/xtt28/neptune/economy/econlookup"
	"github.com/xtt28/neptune/game"
	"github.com/xtt28/neptune/moderation"
	"github.com/xtt28/neptune/moderation/punishments"
	"github.com/xtt28/neptune/scoreboard"
	"github.com/xtt28/neptune/stats"
)

func handleJoin(p *player.Player) {
	activePuns := *punishments.GetActive(p, moderation.PunishmentTypeBan)
	if len(activePuns) > 0 {
		time.AfterFunc(35 * time.Millisecond, func() {
			p.Disconnect(punishments.GenerateMessage(*activePuns[0]))
		})
	}

	p.EnableInstantRespawn()
	scoreboard.Render(p, stats.GetStats(p), econlookup.GetBitsBalance(p))

	game.SendToSpawn(p, false)

	profileTarget := &model.UserProfile{}
	err := database.DB.FirstOrCreate(profileTarget, &model.UserProfile{UUID: p.UUID()}).Error
	if err != nil {
		log.Printf("DB could not load/create user profile for %s (%s): %s", p.Name(), p.UUID(), err.Error())
		p.Disconnect("We could not initialize your user profile. Please try again or contact an administrator.")
		return
	}
	username := strings.ToLower(p.Name())
	if profileTarget.LastUsername != username {
		profileTarget.LastUsername = username
		database.DB.Save(profileTarget)
	}

	p.SendTitle(title.New(text.Colourf("<diamond><bold>Welcome</bold></diamond>")).WithSubtitle(text.Colourf("<grey>to Neptune</grey>")))
}
