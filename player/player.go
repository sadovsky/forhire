package player

import (
	. "github.com/sadovsky/forhire/Character"
)

type Player struct {
	Character
	Xp    int
	Level int
}
