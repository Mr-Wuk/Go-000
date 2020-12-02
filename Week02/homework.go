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
