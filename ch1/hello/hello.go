package main

import (
	"fmt"
)

func main() {
	s := []byte{'0','1','2','3','4','5','6'}
	t := s[1:3]
	t1 := t
	t[0] = 'r'
	t[1] = 'w'
	fmt.Println(string(s))
	fmt.Println(string(t))
	fmt.Println(string(t1))
}