package game

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/npc"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/xtt28/neptune/config"
	"github.com/xtt28/neptune/form"
)

func SpawnDealer(world *world.World) {
	pos := config.NConfig.Locations.DealerCoords
	x := pos[0]
	y := pos[1]
	z := pos[2]

	skin := npc.MustSkin(npc.MustParseTexture("skins/dealer.png"), npc.DefaultModel)

	settings := npc.Settings{
		Name: text.Colourf("<dark-grey><obfuscated>$$</obfuscated></dark-grey> <lapis><bold>THE DEALER</bold></lapis> <dark-grey><obfuscated>$$</obfuscated></dark-grey>\n<gold>Gear and perks</gold>"),
		Scale: 1,
		Immobile: true,
		MainHand: item.NewStack(item.GoldIngot{}, 1),
		Position: mgl64.Vec3{x, y, z},
		Skin: skin,
	}

	npc.Create(settings, world, func(p *player.Player) {
		form.ShowDealerIntermediate(p)
	})
}