package permissions

type PermissionLevel uint

const (
	Default PermissionLevel = 100
	Plus PermissionLevel = 150
	Moderator PermissionLevel = 250
	Admin PermissionLevel = 300
	Owner PermissionLevel = 350
)