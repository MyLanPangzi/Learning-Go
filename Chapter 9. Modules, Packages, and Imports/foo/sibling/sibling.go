package sibling

import "ch9/foo/internal"

func SiblingDoubler(a int) int {
	return internal.Doubler(a)
}
