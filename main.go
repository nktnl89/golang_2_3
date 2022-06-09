package golang_2_3

import "fmt"

// Echo returns message twice
func Echo(message string) string {
	return fmt.Sprintf("%s, - %s", message, message)
}
