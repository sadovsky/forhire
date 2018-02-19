package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/sadovsky/forhire/Character"
	. "github.com/sadovsky/forhire/Player"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	player := new(Player)
	player.Character.Name = "Runner"
	player.Character.Hp = 100
	player.Character.MaxHp = 100
	player.Character.Str = 10
	player.Level = 1

	fmt.Println(player.Character.Hp, ">")

	for i := 0; i < 3; i++ {
		drone := Drone
		Battle(&player.Character, &drone)
	}

	fmt.Println(player.Character.Hp, ">")
}
