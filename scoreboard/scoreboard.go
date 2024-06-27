package scoreboard

import (
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/database/model"
)

func Render(p *player.Player, stats *model.PvPStat, balance uint64) {
	s := scoreboard.New(text.Colourf("<diamond><bold>NEPTUNE</bold></diamond>"))
	s.Set(0, text.Colourf("<grey>%s</grey>", time.Now().Format("02.01.2006")))
	s.Set(1, text.Colourf("<white>  </white>"))
	s.Set(2, text.Colourf("<white><bold> Stats</bold></white>"))
	s.Set(3, text.Colourf("<dark-grey> » </dark-grey><diamond>Kills</diamond> <white>%d</white>", stats.Kills))
	s.Set(4, text.Colourf("<dark-grey> » </dark-grey><diamond>Deaths</diamond> <white>%d</white>", stats.Deaths))
	s.Set(5, text.Colourf("<dark-grey> » </dark-grey><diamond>Bits</diamond> <white>%d</white>", balance))
	s.Set(6, text.Colourf("<grey></grey>"))
	s.Set(7, text.Colourf("<diamond>   serverip.here</diamond>"))

	p.SendScoreboard(s)
}
