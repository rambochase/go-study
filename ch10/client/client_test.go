package client

import (
	"testing"
	"fmt"
	"ch10/reusing"
)

func TestPackage(t *testing.T) {
	reusing.Speak()
	n := 100
	//调用我们自己的包
	if v,err := reusing.GetFib(n); err == nil {
			fmt.Printf("%[1]T,%[1]v",v)
	} else {
		t.Error(err)
	}
}