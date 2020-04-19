package string_test

import (
	"testing"
	"strings"
	"strconv"
)

func TestStringFun(t *testing.T) {
	s := "a,b,c";
	parts := strings.Split(s,",")//切割
	for _,c := range parts {
		t.Log(c)
	}
	t.Log(strings.Join(parts,"-"))//连接

	//整型转字符串
	n := 10
	ns := strconv.Itoa(n)
	t.Log("asdf"+ns)

	//字符串转整型
	if i,err := strconv.Atoi(ns);err == nil {
		t.Log(10+i)
	} else {
		t.Log("it isn't not number string\n",err)
	}
}