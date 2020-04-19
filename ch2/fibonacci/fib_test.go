package fibonacci

import "testing"

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	var tmp int
	t.Log(a," ")
	for i := 0; i<=5; i++ {
		t.Log(b," ")
		tmp = a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 5
	b := 6
	t.Log(a," ",b)
	a,b = b,a
	t.Log(a," ",b)
}