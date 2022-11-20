package jisx4061_test

import (
	"fmt"

	"github.com/shogo82148/jisx4061"
)

func ExampleSort() {
	list := []string{
		"さどう",
		"さとうや",
		"サトー",
		"さと",
		"さど",
		"さとう",
		"さとおや",
	}
	jisx4061.Sort(list)
	for _, s := range list {
		fmt.Println(s)
	}
	// Output:
	// さと
	// さど
	// さとう
	// さどう
	// さとうや
	// サトー
	// さとおや
}
