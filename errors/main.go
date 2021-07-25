package main

import (
	"errors/dao"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	_, err := dao.GetUserInfo(100000000)
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
}
