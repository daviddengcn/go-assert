package assert

import (
	"testing"
)

func TestBasic(t *testing.T) {
	IsTrue(t, "return value", Equals(t, "v", 1, 1))
	IsTrue(t, "return value", NotEquals(t, "v", 1, 4))
	IsTrue(t, "return value", IsTrue(t, "bool", true))
	IsTrue(t, "return value", IsFalse(t, "bool", false))
	IsTrue(t, "return value", StringEquals(t, "string", 1, "1"))
	IsTrue(t, "return value", StrSetEquals(t, "strset", nil, nil))
	IsTrue(t, "return value", LinesEqual(t, "lines", []string{"abc"}, []string{"abc"}))
	
	IsTrue(t, "return value", Maps(t, "appendA", []interface{}{"bcd"}, []interface{}{"bcdA"}, func(src interface{}) interface{} {
		return src.(string) + "A"
	}))
}

func TestShowText(t *testing.T) {
	IsTrue(t, "return value", Equals(t, "v", showText("a\n\r\f\tb"), "a^J^M^L^Ib."))
}
