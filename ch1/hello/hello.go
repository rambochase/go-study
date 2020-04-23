package main

import (
	"os"
	"net/http"
	"fmt"
)

func main() {
	url := "http://www.xinhuanet.com/politics/leaders/2020-04/23/c_1125896815.htm"
	resp,err := http.Get(url)
	if err == nil {
		fmt.Println(resp.Body)
		os.Exit(0)
	}
	fmt.Println(err)
}