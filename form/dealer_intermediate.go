package form

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

func ShowDealerIntermediate(p *player.Player) {
	options := []form.Button{
		form.NewButton("Choose a kit", "textures/items/netherite_sword"),
		form.NewButton("Perk shop", "textures/items/nether_star"),
	}
	f := form.
		NewMenu(DealerIntermediateMenuHandler{}, "The Dealer").
		WithBody("Select an option below to start your journey.").
		WithButtons(options...)
	p.SendForm(f)
}

type DealerIntermediateMenuHandler struct{}

func (h DealerIntermediateMenuHandler) Submit(submitter form.Submitter, pressed form.Button) {
	p, ok := submitter.(*player.Player)
	if !ok {
		return
	}

	if pressed.Text == "Choose a kit" {
		ShowKitSelector(p)
		return
	} else if pressed.Text == "Perk shop" {
		p.Message(text.Colourf("<red><bold>UNSUPPORTED</bold></red> <grey>This feature hasn't been implemented yet.</grey>"))
	}
}
