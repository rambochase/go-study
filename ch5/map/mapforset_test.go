package map_test

import "testing"

func TestMapForSet(t *testing.T) {
	my_set := map[int]bool{}
	my_set[1] = true;
	n := 1
	//判断集合是否存在某个元素
	if my_set[n] {
		t.Logf("%d is in set",n)
	} else {
		t.Logf("%d isn't in set",n)
	}
	
	//获取集合的长度
	len := len(my_set)
	t.Logf("len of set is %d",len)

	n = 2
	my_set[n] = true
	t.Log(my_set)
	//删除集合的元素
	delete(my_set,n)
	t.Log(my_set)
}