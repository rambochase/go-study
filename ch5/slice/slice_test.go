package slice

import "testing"

func TestSlice(t *testing.T) {
	var s0 []int//和数组不同，它不需要指定长度
	s1 := []int{1,2,3,4}
	t.Log(s0)
	t.Log(s1)
	t.Log(len(s0),cap(s0))//0,0
	t.Log(len(s1),cap(s1))
	s0 = append(s0,2)
	s1 = append(s1,5)
	t.Log(s0)
	t.Log(s1)
	t.Log(len(s0),cap(s0))
	t.Log(len(s1),cap(s1))
}

func TestCompareSlice(t *testing.T) {
	s := []int{1,2,3}
	if s == nil {
		t.Log("s is empty")
	} else {
		t.Log(len(s),cap(s))
	}
	
}