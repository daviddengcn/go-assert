package assert

import (
	"testing"
)

func TestBasic(t *testing.T) {
	Equals(t, "v", 1, 1)
	NotEquals(t, "v", 1, 4)
	IsTrue(t, "bool", true)
	IsFalse(t, "bool", false)
	StringEquals(t, "string", 1, "1")
	StrSetEquals(t, "strset", nil, nil)
	LinesEqual(t, "lines", []string{"abc"}, []string{"abc"})
}

func TestShowText(t *testing.T) {
	Equals(t, "v", showText("a\n\r\f\tb"), "a^J^M^L^Ib.")
}
