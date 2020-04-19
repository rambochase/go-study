package function

import (
	"testing"
)

func SumFunc(ops ...int) int {
	sum := 0
	for _,op := range ops {
		sum += op
	}
	return sum
}

func TestFunc2(t *testing.T) {
	sum1 := SumFunc(1,2,3,4,6)
	sum2 := SumFunc(2,3)
	sum3 := SumFunc()
	t.Log(sum1)
	t.Log(sum2)
	t.Log(sum3)
}