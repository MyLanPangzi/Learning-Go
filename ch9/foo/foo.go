package foo

import "ch9/foo/internal"

func FooDoubler(a int) int {
	return internal.Doubler(a)
}
