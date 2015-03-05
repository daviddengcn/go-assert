/*
	Package assert provides some convinient asserting functions on strings, lines, and string sets for testings.
	
	Return values: true if the assert holds, false otherwise.
*/
package assert

import (
	"fmt"
	"path"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/daviddengcn/go-algs/ed"
	"github.com/daviddengcn/go-villa"
)

// escape unprintable chars
func showText(text string) string {
	var buf villa.ByteSlice
	for _, r := range text {
		if r > 0 && r < 27 {
			buf.WriteString(fmt.Sprintf("^%c", 'A'+r-1))
		} else {
			buf.WriteRune(r)
		}
	}
	buf.WriteRune('.')
	return string(buf)
}

// Default: skip == 0
func assertPos(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d: ", path.Base(file), line)
}

/*
	Equals fails the test and shows error message when act and exp are not
	equal
*/
func Equals(t *testing.T, name string, act, exp interface{}) bool {
	if act != exp {
		t.Errorf("%s%s is expected to be %s, but got %v", assertPos(0), name,
			showText(fmt.Sprint(exp)), showText(fmt.Sprint(act)))
		return false
	}
	return true
}

/*
	NotEquals fails the test and shows error message when act and exp are equal
*/
func NotEquals(t *testing.T, name string, act, exp interface{}) bool {
	if act == exp {
		t.Errorf("%s%s is expected not to be %s, but got %v", assertPos(0), name,
			showText(fmt.Sprint(exp)), showText(fmt.Sprint(act)))
		return false
	}
	return true
}

/*
	IsTrue fails the test and shows error message when exp are not true
*/
func IsTrue(t *testing.T, name string, exp bool) bool {
	if exp != true {
		t.Errorf("%s%s unexpectedly got: %v", assertPos(0), name, showText(fmt.Sprint(exp)))
		return false
	}
	return true
}

/*
	IsFalse fails the test and shows error message when exp are not false
*/
func IsFalse(t *testing.T, name string, exp bool) bool {
	if exp != false {
		t.Errorf("%s%s unexpectedly got: %v", assertPos(0), name, showText(fmt.Sprint(exp)))
		return false
	}
	return true
}

/*
	StringEquals fails the test and shows error message when string forms of
	act and exp are not equal
*/
func StringEquals(t *testing.T, name string, act, exp interface{}) bool {
	actS, expS := fmt.Sprintf("%+v", act), fmt.Sprintf("%+v", exp)
	if actS != expS {
		if len(actS) + len(expS) < 70 {
			t.Errorf("%s%s is expected to be %s, but got %v", assertPos(0), name,
				showText(expS), showText(actS))
			return false
		} else {
			t.Errorf("%s%s is expected to be\n%s, but got\n%v", assertPos(0), name,
				showText(expS), showText(actS))
			return false
		}
	}
	return true
}

/*
	StrSetEquals fails the test and shows error message when act and exp are
	not equal string sets.
*/
func StrSetEquals(t *testing.T, name string, act, exp villa.StrSet) bool {
	if !act.Equals(exp) {
		actEls := act.Elements()
		expEls := exp.Elements()
		
		sort.Strings(actEls)
		sort.Strings(expEls)
		
		return linesEquals(t, name, actEls, expEls)
	}
	return true
}

func linesEquals(t *testing.T, name string, act, exp []string) bool {
	if villa.StringSlice(exp).Equals(act) {
		return true
	}
	title := fmt.Sprintf("%sUnexpected %s: ", assertPos(1), name)
	if len(exp) == len(act) {
		title = fmt.Sprintf("%sboth %d lines", title, len(exp))
	} else {
		title = fmt.Sprintf("%sexp %d, act %d lines", title, len(exp), len(act))
	}
	t.Error(title)
	t.Log("Difference(exp ===  act ### change --- +++)")
	_, matA, matB := ed.EditDistanceFFull(len(exp), len(act),
		func(iA, iB int) int {
			sa, sb := exp[iA], act[iB]
			if sa == sb {
				return 0
			}
			return ed.String(sa, sb) + 1
		}, func(iA int) int {
			return len(exp[iA]) + 1
		}, func(iB int) int {
			return len(act[iB]) + 1
		})
	for i, j := 0, 0; i < len(exp) || j < len(act); {
		switch {
		case j >= len(act) || i < len(exp) && matA[i] < 0:
			t.Logf("=== %3d: %s", i+1, showText(exp[i]))
			i++
		case i >= len(exp) || j < len(act) && matB[j] < 0:
			t.Logf("### %3d: %s", j+1, showText(act[j]))
			j++
		default:
			if exp[i] != act[j] {
				t.Logf("--- %3d: %s", i+1, showText(exp[i]))
				t.Logf("+++ %3d: %s", j+1, showText(act[j]))
			} // else
			i++
			j++
		}
	}
	
	return false
}

/*
	LinesEqual fails the test and shows the error message and line-to-line
	differences of the lines when two slices of strings are not equal
*/
func LinesEqual(t *testing.T, name string, act, exp []string) bool {
	return linesEquals(t, name, act, exp)
}

/*
	TextEquals splits input strings into lines and calls LinesEqual
*/
func TextEquals(t *testing.T, name string, act, exp string) bool {
	if act == exp {
		return true
	}
	acts, exps := strings.Split(act, "\n"), strings.Split(exp, "\n")
	return LinesEqual(t, name, acts, exps)
}


/*
	NoError fails the test if err is not nil
*/
func NoError(t *testing.T, err error) bool {
	if err !=  nil {
		t.Errorf("%s%s", assertPos(0), fmt.Sprintf("%v", err))
		return false
	}
	return true
}

/*
	NoErrorf is similar to NoError with an extra format string. The first
	and only variable in the format string should be %v for error.
*/
func NoErrorf(t *testing.T, fmtStr string, err error) bool {
	if err !=  nil {
		t.Errorf("%s%s", assertPos(0), fmt.Sprintf(fmtStr, err))
		return false
	}
	return true
}

/*
Maps asserts whether a mapping function works as expected.
*/
func Maps(t *testing.T, name string, src []interface{}, dst []interface{}, f func(src interface{}) interface{}) bool {
	succ := true
	for idx, s := range src {
		act := f(s)
		exp := dst[idx]
		if act != exp {
			t.Errorf("%s%s is expected to be mapped by %s as %s, but got %v", assertPos(0), showText(fmt.Sprint(s)), name,
				showText(fmt.Sprint(exp)), showText(fmt.Sprint(act)))
			succ = false
		}
	}
	
	return succ
}
