package reusing

import (
	"errors"
)

//字母开头大写才能被包外访问
func GetFib(len int) ([]int,error) {
	if len < 2 || len > 100 {
		return nil,errors.New("len should be in [2,100]")
	}
	list := []int{1,1}
	for i := 2; i < len; i++ {
		list = append(list,list[i-1] + list[i-2])
	}
	return list,nil
}
