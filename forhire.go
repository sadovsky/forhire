package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/sadovsky/forhire/character"
	. "github.com/sadovsky/forhire/game"
	. "github.com/sadovsky/forhire/player"
	. "github.com/sadovsky/forhire/term"
)

func gameLoop() {

	//initialize
	rand.Seed(time.Now().UTC().UnixNano())

	game := new(Game)
	player := new(Player)

	//TODO ask user
	if game.SaveExists() {
		fmt.Print("Load saved game? (Y/N)")
		input := GetCh()
		switch input {
		case "y", "Y":
			player = game.Load()
			fmt.Println()
			fmt.Println("Game Loaded!")
			break
		case "n", "N":
			player.Character.Name = "Runner"
			player.Initialize()
			break
		}
	} else {
		player.Character.Name = "Runner"
		player.Initialize()
	}

GameLoop:
	for {

		if player.IsDead() {
			ClearScreen()
			fmt.Println("+-------------------------------+")
			fmt.Println(player.Character.Name, "has perished...")
			fmt.Println("+-------------------------------+")
			break GameLoop
		}

		ClearScreen()
		fmt.Println("+-------------------------------+")
		fmt.Println("B to (B)attle")
		fmt.Println("I for (I)nventory")
		fmt.Println("S to (S)ave")
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
			player.Battle(&drone)
			GetCh()
			continue
		case "s", "S":
			fmt.Print(input)
			fmt.Println()
			game.Save(player)
			break GameLoop
		case "i", "I":
			fmt.Print(input)
			fmt.Println()
			player.PrintInventory()
			GetCh()
			continue
		}

		fmt.Println()
	}
}

func main() {
	gameLoop()
}
