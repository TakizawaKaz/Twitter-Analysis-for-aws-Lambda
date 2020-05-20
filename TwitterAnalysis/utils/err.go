package utils

import (
	"fmt"
)

func ErrorRaise(m string) error {
	return fmt.Errorf(m)
}
