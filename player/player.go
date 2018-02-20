package player

import (
	"fmt"

	. "github.com/sadovsky/forhire/character"
	. "github.com/sadovsky/forhire/item"
)

type Player struct {
	Character
	Xp        int
	Level     int
	Inventory []Item
	Credits   int
}

func (p Player) Prompt() {
	fmt.Print("HP:", p.Character.Hp, "> ")
}

func (p Player) IsDead() bool {
	if p.Character.Hp <= 0 {
		return true
	}

	return false
}

func (p *Player) Initialize() {
	p.Character.Hp = 100
	p.Character.MaxHp = 100
	p.Character.Str = 20
	p.Level = 1
	p.Inventory = append(p.Inventory, LeatherJacket)
	p.Credits = 125
	p.Character.Description = "The neophyte player"
}

func (p Player) PrintInventory() {
	fmt.Println("Credits (₿)", ":", p.Credits)
	for i, v := range p.Inventory {
		fmt.Println(i, ":", v.Name)
	}
}
