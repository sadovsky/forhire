package character

import (
	"fmt"
	"math/rand"
)

type Character struct {
	Name        string
	Description string
	Str         int
	Hp          int
	MaxHp       int
	Xp          int
}

func Battle(a *Character, b *Character) {
	var at1, at2 int
	println()
	println(a.Description, "encounters", b.Description)

	for (a.Hp > 0) && (b.Hp > 0) {
		fmt.Println()
		fmt.Println("<", a.Name, " HP: ", a.Hp, " VS ", b.Name, " HP: ", b.Hp, ">")

		at1 = rand.Intn(a.Str)
		at2 = rand.Intn(b.Str)

		b.Hp -= at1
		fmt.Println(a.Name, " attacks ", b.Name, " for ", at1, "damage!")

		if b.Hp <= 0 {
			fmt.Println(a.Name, " slays ", b.Name)
		} else {
			a.Hp -= at2
			fmt.Println(b.Name, " attacks ", a.Name, " for ", at2, "damage!")
			if a.Hp <= 0 {
				fmt.Println(b.Name, " slays ", a.Name)
			}
		}
	}

}
