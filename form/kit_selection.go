package form

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/xtt28/neptune/game/kit"
)

func ShowKitSelector(p *player.Player) {
	kitButtons := []form.Button{}
	for _, kit := range kit.Kits {
		kitButtons = append(kitButtons, kit.ToButton())
	}
	f := form.
		NewMenu(KitSelectorMenuHandler{}, "Choose a kit").
		WithButtons(kitButtons...)
	p.SendForm(f)
}

type KitSelectorMenuHandler struct{}

func (h KitSelectorMenuHandler) Submit(submitter form.Submitter, pressed form.Button) {
	p, ok := submitter.(*player.Player)
	if !ok {
		return
	}

	kitButtons := []form.Button{}
	for _, kit := range kit.Kits {
		kitButtons = append(kitButtons, kit.ToButton())
	}

	index := -1

	for i, v := range kitButtons {
		if v == pressed {
			index = i
			break
		}
	}

	kit.Kits[index].GiveTo(p)
}
