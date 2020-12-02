### 作业

----

> 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
> 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

### 解题思路

> 在dao层出现错误后，要向上抛出，在顶层只获取错误根本原因。

### 代码

1. homework.go

   ```go
   package week02_error
   
   import (
   	_errors "github.com/pkg/errors"
   )
   
   //我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
   //是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
   
   //应该抛给上层，具有最高可从用性的包只能返回根错误值
   func Dao() error {
   	return _errors.Wrap(_errors.New("sql.ErrNoRows"), "dao sql.ErrNoRows")
   }
   func Service() error {
   	err := _errors.WithMessage(Dao(), "service err")
   	return err
   }
   ```

   

2. homework_test.go

   ```go
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
   ```

