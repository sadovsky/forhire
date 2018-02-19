package main

import (
	"fmt"
	"math/rand"
	"time"
)

type character struct {
	name string
	str  int
	hp   int
}

func Battle(a *character, b *character) {
	var at1 int
	var at2 int
	for (a.hp > 0) && (b.hp > 0) {
		at1 = rand.Intn(a.str)
		at2 = rand.Intn(b.str)
		a.hp -= at2
		fmt.Println(a.hp, "> ", a.name, " attacks ", b.name, " for ", at1, "damage!")
		b.hp -= at1
		fmt.Println(a.hp, ">", b.name, " attacks ", a.name, " for ", at2, "damage!")
	}
	if a.hp > 0 {
		fmt.Println(a.hp, ">", a.name, " slays ", b.name)
	}
	if b.hp > 0 {
		fmt.Println(a.hp, ">", b.name, " slays ", a.name)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	tank := &character{"Runner", 20, 100}

	fmt.Println(tank.hp, ">")

	for i := 0; i < 5; i++ {
		drone := &character{"Drone", 5, 25}
		Battle(tank, drone)
	}

	fmt.Println(tank.hp, ">")
}
