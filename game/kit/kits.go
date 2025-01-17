package kit

import (
	"math"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
)

var Kits = []Kit{
	{
		Name:        "Soldier",
		Description: "Melee expert with superior combat skills.",
		Icon:        "textures/items/netherite_sword",
		Helmet: item.NewStack(item.Helmet{Tier: item.ArmourTierGold{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierIron{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Boots: item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 2)),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierNetherite}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Sharpness{}, 2)),
		},
	},
	{
		Name:        "Archer",
		Description: "Ranged specialist with precise attacks.",
		Icon:        "textures/items/bow_standby",
		Helmet: item.NewStack(item.Helmet{Tier: item.ArmourTierDiamond{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierChain{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.ProjectileProtection{}, 4)),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 2)),
		Boots: item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierIron}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Sharpness{}, 1)),
			item.NewStack(item.Bow{}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Power{}, 3)),
			item.NewStack(item.Arrow{}, 32),
		},
	},
	{
		Name:        "Tank",
		Description: "Durable tank with high damage resistance.",
		Icon:        "textures/items/netherite_chestplate",
		Helmet: item.NewStack(item.Helmet{Tier: item.ArmourTierGold{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 3)),
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierNetherite{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 2)),
		Boots: item.NewStack(item.Boots{Tier: item.ArmourTierDiamond{}}, 1),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierNetherite}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Sharpness{}, 1)),
		},
		Effects: []effect.Effect{
			effect.New(effect.Slowness{}, 1, math.MaxInt64),
		},
	},
	{
		Name:        "Jupiter",
		Description: "Lightning wielder with powerful strikes.",
		Icon:        "textures/blocks/lightning_rod",
		Helmet: item.NewStack(item.Helmet{Tier: item.ArmourTierIron{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierChain{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 3)),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1),
		Boots:    item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierGold}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Sharpness{}, 3)).
				WithValue(ItemAbilityKey, JupiterAbilityValue),
		},
	},
	{
		Name:        "Ninja",
		Description: "Fast assassin with unmatched speed.",
		Icon:        "textures/items/nether_star",
		Helmet: item.NewStack(item.Helmet{Tier: item.ArmourTierIron{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 1)),
		Chestplate: item.NewStack(item.Chestplate{Tier: item.ArmourTierDiamond{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.ProjectileProtection{}, 2)),
		Leggings: item.NewStack(item.Leggings{Tier: item.ArmourTierIron{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 3)),
		Boots: item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1).
			WithEnchantments(item.NewEnchantment(enchantment.Protection{}, 2)),
		Items: []item.Stack{
			item.NewStack(item.Sword{Tier: item.ToolTierGold}, 1).
				WithEnchantments(item.NewEnchantment(enchantment.Sharpness{}, 3)),
		},
		Effects: []effect.Effect{
			effect.New(effect.Speed{}, 1, math.MaxInt64),
		},
	},
}

type Kit struct {
	Name        string
	Description string
	Icon        string
	Helmet      item.Stack
	Chestplate  item.Stack
	Leggings    item.Stack
	Boots       item.Stack
	Items       []item.Stack
	Effects     []effect.Effect
}

func (k Kit) GiveTo(target *player.Player) {
	target.Inventory().Clear()
	target.Armour().Clear()

	for i, v := range k.Items {
		target.Inventory().SetItem(i, v)
	}
	for _, v := range k.Effects {
		target.AddEffect(v)
	}

	target.Armour().Set(k.Helmet, k.Chestplate, k.Leggings, k.Boots)
}

func (k Kit) ToButton() form.Button {
	return form.NewButton(k.Name, k.Icon)
}
