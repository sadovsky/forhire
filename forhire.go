package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/sadovsky/forhire/character"
	. "github.com/sadovsky/forhire/player"
	. "github.com/sadovsky/forhire/term"
)

func gameLoop() {

	//initialize
	rand.Seed(time.Now().UTC().UnixNano())

	player := new(Player)
	player.Character.Name = "Runner"
	player.Initialize()

	//check for death

GameLoop:
	for {
		ClearScreen()
		fmt.Println("+-------------------------------+")
		fmt.Println("B to (B)attle")
		fmt.Println("Q to (Q)uit")
		fmt.Println("+-------------------------------+")
		player.Prompt()

		input := GetCh()
		switch input {
		case "q", "Q":
			fmt.Print(input)
			fmt.Println()
			break GameLoop
		case "b", "B":
			fmt.Print(input)
			drone := Drone
			Battle(&player.Character, &drone)
			GetCh()
			continue
		}

		fmt.Println()
	}
}

func main() {
	gameLoop()
}
