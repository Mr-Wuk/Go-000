package week02_error_test

import (
	"fmt"
	_errors "github.com/pkg/errors"
	week02_error "practice/src/week02-error"
	"testing"
)

func TestHomeWork(t *testing.T) {
	err := week02_error.Service()
	//获取根错误值
	err = _errors.Cause(err)
	fmt.Printf("%+v\n", err)
}
