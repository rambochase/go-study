package array

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1,2,3}
	arr2 := [...]int{1,2,3}
	t.Log(arr[0],arr[1])//会初始化为0值
	t.Log(arr1)
	t.Log(arr2)
}

//数组遍历
func TestArrayTravel(t *testing.T) {
	arr := [...]int{5,7,1,2,3,2,0,1,7,8,4,6,9}

	len := len(arr)
	for i:=0; i < len; i++ {
		t.Log(arr[i])
	}

	//foreach
	for idx,e := range arr {
		t.Log(idx,e)
	}

	for _,e := range arr {
		t.Log(e)
	}
}