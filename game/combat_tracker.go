package game

import (
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/player"
)

var Combat = &CombatTracker{Combats: make(map[*player.Player]*CombatEntry)}

type CombatTracker struct {
	mu      sync.Mutex
	Combats map[*player.Player]*CombatEntry
}

type CombatEntry struct {
	Attacker *player.Player
	Timer    *time.Timer
}

func (ct *CombatTracker) RecordHit(victim, attacker *player.Player) {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	if entry, exists := ct.Combats[victim]; exists && entry.Attacker != attacker {
		entry.Timer.Stop()
		delete(ct.Combats, victim)
	}

	if entry, exists := ct.Combats[victim]; exists {
		entry.Timer.Stop()
		entry.Timer = time.AfterFunc(10*time.Second, func() {
			ct.Clear(victim)
		})
	} else {
		ct.Combats[victim] = &CombatEntry{
			Attacker: attacker,
			Timer: time.AfterFunc(10*time.Second, func() {
				ct.Clear(victim)
			}),
		}
	}
}

func (ct *CombatTracker) Clear(victim *player.Player) {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	if entry, exists := ct.Combats[victim]; exists {
		entry.Timer.Stop()
		delete(ct.Combats, victim)
	}
}
