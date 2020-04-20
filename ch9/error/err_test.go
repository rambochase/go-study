package error_test

import (
	"testing"
	"fmt"
	"errors"
)

//获取斐波那契数列
func getFibonacci(len int) ([]int,error) {
	if len < 2 || len > 100 {
		return nil,errors.New("len should be in [2,100]")
	}
	list := []int{1,1}
	for i := 2; i < len; i++ {
		next := list[i-1] + list[i-2]
		list = append(list,next)
	}
	return list,nil
}

func TestError(t *testing.T) {
	len := 1
	if list,err := getFibonacci(len);err == nil {
		fmt.Printf("%[1]T\n%[1]v\n",list)
	} else {
		t.Error(err)
		//fmt.Printf("%s\n",err)
	}	
}