package condition

import "testing"

func TestSwitch(t *testing.T) {
	a := 10
	switch a {
		case 0: t.Log(0)
		case 1: t.Log(1)
		case 2: t.Log(2)
		default: t.Log("default")
	}

	//相当if
	switch {
		case a < 10: t.Log("不完美")
		case a == 10: t.Log("完美")
		case a > 10: t.Log("超级完美")
	}
	
	//switch的case支持多个项
	for i := 0; i < 5; i++ {
		switch i {
		case 0,2: t.Log("偶数")
		case 1,3: t.Log("奇数")
		default: t.Log("不在0-3")
		}
	}
}