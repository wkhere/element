package element

import (
	"fmt"
)

var Die func(a ...interface{})

func init() {
	Die = die
}

func die(a ...interface{}) {
	panic(fmt.Sprint(a...))
}
