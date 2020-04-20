package error

import (
	//"os"
	"testing"
	"errors"
	"fmt"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		t.Log("finnally")
	}()
	fmt.Println("start")
	panic(errors.New("there is a error"))
	//os.Exit(-1)
}