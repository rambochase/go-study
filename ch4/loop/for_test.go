package loop

import "testing"

func TestWhile(t *testing.T) {
	//while条件循环
	//while (n<5)
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}
}

func TestWhileTrue(t *testing.T) {
	//死循环
	n := 0
	for {
		t.Log(n)
		n++
		if n > 20 {
			break;
		}
	}
}