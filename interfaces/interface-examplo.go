package main

import (
	"strconv"
)

/*
 * Interfaces examples
 *
 * References by:
 * https://research.swtch.com/interfaces
 */
type(
	Stringer interface {
		String() string
	}
)

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}

	switch v := any.(type) {
	case int64:
		s := strconv.FormatInt(v, 64)
		return s
	case float64:
		f := strconv.FormatFloat(v, 'E', -1, 64)
		return f
	}
	return "???"
}

func main() {
	var i interface{}
	ToString(i)
}