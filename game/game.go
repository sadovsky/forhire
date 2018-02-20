package game

import (
	"encoding/gob"
	"fmt"
	"os"

	. "github.com/sadovsky/forhire/player"
)

type Game struct {
	Stage int
}

func (g Game) SaveExists() bool {
	return exists("player.gob")
}

func (g Game) Save(p *Player) {
	err := writeGob("player.gob", p)
	if err != nil {
		fmt.Println(err)
	}
}

func (g Game) Load() *Player {
	var p = new(Player)
	err := readGob("player.gob", p)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func exists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	} else {
		panic(err)
	}
	file.Close()
	return err
}
