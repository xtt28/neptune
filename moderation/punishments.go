package moderation

type PunishmentType string

const (
	PunishmentTypeBan  PunishmentType = "ban"
	PunishmentTypeMute PunishmentType = "mute"
	PunishmentTypeKick PunishmentType = "kick"
)
