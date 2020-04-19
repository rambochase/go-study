package main

import (
	"fmt"
	"os"
)

func main() {
	array := []int{5,4,3,2,1}//切片
	fmt.Println(array)
	BubbleSort(array)
	fmt.Println(array)
	os.Exit(0)
}

//冒泡排序
func BubbleSort(pointer []int) {
	arrayLen := len(pointer)
	for i := 0; i < arrayLen; i++ {
		for j := 0; j < arrayLen - i - 1; j++ {
			if pointer[j] > pointer[j+1] {
				tmp := pointer[j]
				pointer[j] = pointer[j+1]
				pointer[j+1] = tmp
			}
		}
	}
}