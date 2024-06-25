package handlers

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/xtt28/neptune/config"
	"github.com/xtt28/neptune/game"
)

func (m *BasePlayerHandler) HandleRespawn(pos *mgl64.Vec3, w **world.World) {
	spawn := config.NConfig.Locations.Spawn
	*pos = mgl64.Vec3{spawn[0], spawn[1], spawn[2]}
	game.SendToSpawn(m.p, true)
}