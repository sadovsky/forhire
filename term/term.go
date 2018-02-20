package term

import (
	"fmt"

	"github.com/pkg/term"
)

func GetCh() string {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)
	t.Restore()
	t.Close()
	if err != nil {
		panic(err)
	}
	return string(bytes[0:numRead])
}

func ClearScreen() {
	fmt.Printf("[H[J")
}
