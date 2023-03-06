package returnResult

import (
	logger "gingo/src/common/logger"
)

const (
	success string = "success"
	fail    string = "failure"
)

type Result struct {
	Status  string
	Message string
	Object  interface{}
}

type Grid struct {
	Total int
	Data  interface{}
}

func Success(msg string) Result {
	return Result{Status: success, Message: msg}
}

func SuccessObject(msg string, obj interface{}) Result {
	return Result{Status: success, Message: msg, Object: obj}
}

func SuccessGrid(total int, data interface{}) Result {
	return Result{Status: success, Message: "操作成功", Object: Grid{Total: total, Data: data}}
}

func Fail(msg string) Result {
	logger.Error(msg)
	return Result{Status: fail, Message: msg}
}

func FailWarn(msg string) Result {
	logger.Warn(msg)
	return Result{Status: fail, Message: msg}
}
func FailNoError(msg string) Result {
	logger.Info(msg)
	return Result{Status: fail, Message: msg}
}

func FailObject(msg string, object interface{}) Result {
	logger.Error()
	return Result{Status: fail, Message: msg, Object: object}
}

func FailCustomer(msg string, object interface{}) Result {
	return Result{Status: fail, Message: msg, Object: object}
}
