package player

import (
	"fmt"

	. "github.com/sadovsky/forhire/character"
)

type Player struct {
	Character
	Xp    int
	Level int
}

func (p Player) Prompt() {
	fmt.Print("HP:", p.Character.Hp, "> ")
}

func (p *Player) Initialize() {
	p.Character.Hp = 100
	p.Character.MaxHp = 100
	p.Character.Str = 20
	p.Level = 1
}
