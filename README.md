go-test [![GoSearch](http://go-search.org/badge?id=github.com%2Fdaviddengcn%2Fgo-assert)](http://go-search.org/view?id=github.com%2Fdaviddengcn%2Fgo-assert)
=======

Assert utils for GO testing.

Featured
--------
* [LinesEqual](http://godoc.org/github.com/daviddengcn/go-assert#LinesEqual) and [TextEquals](http://godoc.org/github.com/daviddengcn/go-assert#TextEquals)
Other than showing the assertion failure message, they also shows the line-to-line diff result for your information:

**Example code**
```go
func TestForExample(t *testing.T) {
	TextEquals(t, "info",
`Hello world,
Pleae help me,
doing this`,
`Hello world,
please help me`)
}
```

**Results**
```
--- FAIL: TestForExample (0.00 seconds)
	assert.go:123: assert.go:172: Unexpected info: exp 2, act 3 lines
	assert.go:124: Difference(exp ===  act ### change --- +++)
	assert.go:147: ---   2: please help me.
	assert.go:148: +++   2: Pleae help me,.
	assert.go:143: ###   3: doing this.
```

License
-------
BSD license.
