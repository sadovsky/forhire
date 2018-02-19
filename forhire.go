package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/sadovsky/forhire/Character"
)

type Player struct {
	Character
	xp    int
	level int
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	player := new(Player)
	player.Character.Name = "Runner"
	player.Character.Hp = 100
	player.Character.Str = 10
	player.level = 1

	fmt.Println(player.Character.Hp, ">")

	for i := 0; i < 3; i++ {
		drone := Drone
		//drone := &Character{"Drone", 5, 25}
		Battle(&player.Character, &drone)
	}

	fmt.Println(player.Character.Hp, ">")
}
