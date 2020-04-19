package map_test

import "testing"

func TestMap(t *testing.T) {
	m0 := map[string]int{}
	m1 := map[string]int{"one":1,"two":2}
	m2 := make(map[string]int,10)

	m0["zero"] = 0
	t.Log(len(m0))
	t.Log(m0)
	t.Log(m1)
	t.Log(m2)
	t.Log(len(m2))
}

//判断key是否存在
func TestMapExistKey(t *testing.T) {
	m := map[string]string{"one":"1","two":"2"}
	t.Log(m["three"])//不会报错，返回0
	if value , ok := m["three"]; ok {
		t.Logf("m[\"three\"] is exist,value is %s",value)
	} else {
		t.Log("m[\"three\"] isn't exist")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[string]int{"zero":0,"one":1,"two":2}
	for k,v := range m {
		t.Logf("index=%s,value=%d",k,v)
	}

	for _,v := range m {
		t.Logf("value=%d",v)
	}
}