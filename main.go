package golang_2_3

import (
	"fmt"
	"math/rand"
)

// Echo returns message and random int
func Echo(message string) string {
	return fmt.Sprintf("%s, - %d", message, rand.Int())
}
