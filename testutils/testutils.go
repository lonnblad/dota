package testutils

import (
	"fmt"
	"testing"
)

func Assert(x, y interface{}, t *testing.T) {
	if x == nil && x == y {
		return
	} else if y == nil {
		return
	} else if x != y {
		t.Errorf("\n===\n%+v\n%+v\n===", x, y)
	}
}

func AssertError(x, y error, t *testing.T) {
	if x == nil && x == y {
		return
	} else if fmt.Sprint(x) != fmt.Sprint(y) {
		t.Errorf("\n===\n%+v\n%+v\n===", x, y)
	}
}
