package character

import (
	"fmt"
	"math/rand"
)

type Character struct {
	Name  string
	Str   int
	Hp    int
	MaxHp int
}

func Battle(a *Character, b *Character) {
	var at1 int
	var at2 int
	for (a.Hp > 0) && (b.Hp > 0) {
		at1 = rand.Intn(a.Str)
		at2 = rand.Intn(b.Str)
		a.Hp -= at2
		fmt.Println(a.Hp, "> ", a.Name, " attacks ", b.Name, " for ", at1, "damage!")
		b.Hp -= at1
		fmt.Println(a.Hp, ">", b.Name, " attacks ", a.Name, " for ", at2, "damage!")
	}
	if a.Hp > 0 {
		fmt.Println(a.Hp, ">", a.Name, " slays ", b.Name)
	}
	if b.Hp > 0 {
		fmt.Println(a.Hp, ">", b.Name, " slays ", a.Name)
	}
}