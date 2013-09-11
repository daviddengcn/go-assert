package assert

import (
	"testing"
)

func TestBasic(t *testing.T) {
	Equals(t, "name", 1, 1)
	StringEquals(t, "name", 1, "1")
	StrSetEquals(t, "name", nil, nil)
}
