package assert

import (
	"fmt"
	"strings"
	"testing"

	"github.com/daviddengcn/go-algs/ed"
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

// escape unprintable chars
func showText(text string) string {
	var buf villa.ByteSlice
	for _, r := range text {
		if r > 0 && r < 27 {
			buf.WriteString(fmt.Sprintf("^%c", 'A' + r - 1))
		} else {
			buf.WriteRune(r)
		}
	}
	buf.WriteRune('.')
	return string(buf)
}

func LinesEqual(t *testing.T, name string, act, exp []string) {
	if villa.StringSlice(exp).Equals(act) {
		return
	}
	title := fmt.Sprintf("Unexpected %s: ", name)
	if len(exp) == len(act) {
		title = fmt.Sprintf("%sboth %d lines", title, len(exp))
	} else {
		title = fmt.Sprintf("%sexp %d, act %d lines", title, len(exp), len(act))
	}
	t.Error(title)
	t.Log("Difference(exp ---  act +++)")
	_, matA, matB := ed.EditDistanceFFull(len(exp), len(act), func(iA, iB int) int {
		sa, sb := exp[iA], act[iB]
		if sa == sb {
			return 0
		}
		return ed.String(sa, sb)
	}, func(iA int) int {
		return len(exp[iA]) + 1
	}, func(iB int) int {
		return len(act[iB]) + 1
	})
	for i, j := 0, 0; i < len(exp) || j < len(act); {
		switch {
		case j >= len(act) || i < len(exp) && matA[i] < 0:
			t.Logf("--- %3d: %s", i+1, showText(exp[i]))
			i++
		case i >= len(exp) || j < len(act) && matB[j] < 0:
			t.Logf("+++ %3d: %s", j+1, showText(act[j]))
			j++
		default:
			if exp[i] != act[j] {
				t.Logf("--- %3d: %s", i+1, showText(exp[i]))
				t.Logf("+++ %3d: %s", j+1, showText(act[j]))
			} // else
			i++
			j++
		}
	} // for i, j
}

func TextEquals(t *testing.T, name string, act, exp string) {
	if act == exp {
		return
	}
	acts, exps := strings.Split(act, "\n"), strings.Split(exp, "\n")
	LinesEqual(t, name, acts, exps)
}