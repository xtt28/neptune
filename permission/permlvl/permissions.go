package permlvl

type PermissionLevel uint

const (
	LvlDefault   PermissionLevel = 100
	LvlPlus      PermissionLevel = 150
	LvlModerator PermissionLevel = 250
	LvlAdmin     PermissionLevel = 300
	LvlOwner     PermissionLevel = 350
)

var LevelToName = map[PermissionLevel]string{
	100: "Default",
	150: "Plus",
	250: "Moderator",
	300: "Admin",
	350: "Owner",
}

var NameToLevel = map[string]PermissionLevel{
	"default":   100,
	"plus":      150,
	"moderator": 250,
	"admin":     300,
	"owner":     350,
}

var LevelToChatFormat = map[PermissionLevel]string{
	100: "<grey>%v <dark-grey>»</dark-grey> %v</grey>",
	150: "<diamond><bold>+</bold> %v</diamond> <dark-grey>»</dark-grey> <quartz>%v</quartz>",
	250: "<emerald><bold>MOD</bold> %v</emerald> <dark-grey>»</dark-grey> <quartz>%v</quartz>",
	300: "<red><bold>ADMIN</bold> %v</red> <dark-grey>»</dark-grey> <quartz>%v</quartz>",
	350: "<redstone><bold>OWNER</bold> %v</redstone> <dark-grey>»</dark-grey> <quartz>%v</quartz>",
}
