package data_type

import "testing"

func TestString(t *testing.T) {
	var str string
	str = "wer"
	strPt := &str //指针
	t.Log(str,strPt)
	t.Logf("%T,%T",str,strPt)//把类型打印出来
}