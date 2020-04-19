package condition

import "testing"

func TestIf(t *testing.T) {
	a := 20
	if a > 60 && a < 70 {
		t.Log("及格")
	} else if a >= 70 && a < 80 {
		t.Log("中等")
	} else if a >= 80 {
		t.Log("优秀")
	} else {
		t.Log("不及格")
	}
}