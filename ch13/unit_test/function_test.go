package unit_test

import (
	"testing"
	"fmt"
)

func square(n int) int {
	return n * n * n
}

func sql() {
	fmt.Println("No Done")
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	len := len(inputs)
	for i := 0; i < len; i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d,but the actual is %d",
		inputs[i],expected[i],ret)
		} else {
			t.Log("success")
		}
	}
}