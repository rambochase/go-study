package error

import (
	"testing"
	"errors"
	"fmt"
)

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover();err != nil {
			fmt.Println("recovered from ",err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("there is a error"))
}