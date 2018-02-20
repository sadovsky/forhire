package player

import (
	"fmt"
	"math/rand"

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

func (p *Player) CheckLevel() {
	//TODO actually make an XP table
	if p.Xp > 10 {
		p.Level++
		p.Str += 3
		fmt.Println(p.Name, "has reached level", p.Level)
		fmt.Println("Strength increased by 3")
	}
}

func (p *Player) Battle(npc *Character) {
	var at1, at2 int
	println()
	println(p.Character.Description, "encounters", npc.Description)

	for (p.Character.Hp > 0) && (npc.Hp > 0) {
		fmt.Println()
		fmt.Println("<", p.Character.Name, " HP: ", p.Character.Hp, " VS ", npc.Name, " HP: ", npc.Hp, ">")

		at1 = rand.Intn(p.Character.Str)
		at2 = rand.Intn(npc.Str)

		npc.Hp -= at1
		fmt.Println(p.Character.Name, " attacks ", npc.Name, " for ", at1, "damage!")

		if npc.Hp <= 0 {
			fmt.Println(p.Character.Name, " slays ", npc.Name)
			//add up xp
			p.Xp += npc.Xp
			p.CheckLevel()
		} else {
			p.Character.Hp -= at2
			fmt.Println(npc.Name, " attacks ", p.Character.Name, " for ", at2, "damage!")
			if p.IsDead() {
				fmt.Println(npc.Name, " slays ", p.Character.Name)
			}
		}
	}

}
