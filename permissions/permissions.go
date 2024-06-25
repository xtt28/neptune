package permissions

type PermissionLevel uint

const (
	LvlDefault   PermissionLevel = 100
	LvlPlus      PermissionLevel = 150
	LvlModerator PermissionLevel = 250
	LvlAdmin     PermissionLevel = 300
	LvlOwner     PermissionLevel = 350
)
