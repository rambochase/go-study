package string_test

import "testing"

func TestTravelString(t *testing.T) {
	c := "中华人民共和国"
	for _,v := range c {
		t.Logf("%[1]c,UTF-8:%[1]x",v)
	}
	for _,v := range c {
		t.Logf("%c",v)
	}
}