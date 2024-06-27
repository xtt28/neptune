package form

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/stats"
)

func ShowStats(p *player.Player) {
	stats := stats.GetStats(p)

	var kdr float64
	if stats.Deaths != 0 {
		kdr = float64(stats.Kills) / float64(stats.Deaths)
	}

	f := form.
		NewMenu(StatsMenuHandler{}, "Your statistics").
		WithBody(text.Colourf(
			"Kills: <bold>%d</bold>\nDeaths: <bold>%d</bold>\nK/D: <bold>%.2f</bold>",
			stats.Kills,
			stats.Deaths,
			kdr)).
		WithButtons(form.NewButton("Close", ""))
	p.SendForm(f)
}

type StatsMenuHandler struct{}

func (h StatsMenuHandler) Submit(submitter form.Submitter, pressed form.Button) {
	// Do nothing.
}
