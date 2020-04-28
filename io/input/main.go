package main

import (
	"fmt"
)

func main() {
	key := ""
	fmt.Println("输入一个数：")
	fmt.Scanln(&key)
	fmt.Printf("你输入的是：%[1]T,%[1]v\n",key)
}