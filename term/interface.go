package term

import (
	"fmt"
)

func ClearScreen() {
	fmt.Printf("[H[J")
}
