package assert

import (
	"fmt"
	"testing"

	"github.com/daviddengcn/go-villa"
)

/*
	AssertEquals shows error message when act and exp don't equal
*/
func Equals(t *testing.T, name string, act, exp interface{}) {
	if act != exp {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	}
}

/*
	AssertEquals shows error message when strings forms of act and exp don't
	equal
*/
func StringEquals(t *testing.T, name string, act, exp interface{}) {
	if fmt.Sprintf("%v", act) != fmt.Sprintf("%v", exp) {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	} // if
}

/*
	AssertStrSetEquals shows error message when act and exp are equal string
	sets.
*/
func StrSetEquals(t *testing.T, name string, act, exp villa.StrSet) {
	if !act.Equals(exp) {
		t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
	}
}
