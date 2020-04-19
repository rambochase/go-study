package array

import "testing"

//数组截取
func TestInterception(t *testing.T) {
	array := [...]int{0,1,2,3,4,5}
	tmp := array[1:2]//[1]
	tmp1 := array[1:3]//[1,2]
	tmp2 := array[:3]//[0,1,2]
	tmp3 := array[1:]//[1,2,3,4,5]
	tmp4 := array[1:len(array)]//[1,2,3,4,5]
	tmp5 := array[:]//[0,1,2,3,4,5]
	t.Log(tmp)
	t.Log(tmp1)
	t.Log(tmp2)
	t.Log(tmp3)
	t.Log(tmp4)
	t.Log(tmp5)
}