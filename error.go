package element

import (
	"fmt"
)

var Die func(format string, a ...interface{})

func init() {
	Die = die
}

func die(format string, a ...interface{}) {
	panic(fmt.Sprint(a...))
}
