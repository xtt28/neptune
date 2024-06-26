package game

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

var Kits = []Kit{
	{
		Name: "Soldier",
		Description: "Best for melee combat",
		Icon: "textures/items/netherite_sword",
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierIron{}}, 1),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierNetherite}, 1),
		},
	},
	{
		Name: "Archer",
		Description: "Best for ranged combat",
		Icon: "textures/items/bow_standby",
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierIron{}}, 1),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierIron}, 1),
			item.NewStack(item.Bow{}, 1),
			item.NewStack(item.Arrow{}, 32),
		},
	},
}

type Kit struct {
	Name string
	Description string
	Icon string
	Helmet item.Stack
	Chestplate item.Stack
	Leggings item.Stack
	Boots item.Stack
	Items []item.Stack
}

func (k Kit) GiveTo(target *player.Player) {
	target.Inventory().Clear()
	target.Armour().Clear()

	for i, v := range k.Items {
		target.Inventory().SetItem(i, v)
	}
	
	target.Armour().Set(k.Helmet, k.Chestplate, k.Leggings, k.Boots)
}

func (k Kit) ToButton() form.Button {
	return form.NewButton(text.Colourf("<bold>%s</bold>\n%s", k.Name, k.Description), k.Icon)
}