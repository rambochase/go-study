package main

import (
	"os"
	"fmt"
	"strings"
)

//go run args.go hello
func main() {
	//os.Args是一个切片
	for _,v := range os.Args[1:] {
		fmt.Println(v)
	}
	s := strings.Join(os.Args[1:],",")
	fmt.Println(s)
}