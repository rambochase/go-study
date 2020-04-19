package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int {1,2,3,4,5}
	b := [...]int {2,3,4,5,6}
	c := [...]int {1,2,3,5,4}
	t.Log(a == b)
	t.Log(a == c)
}