package tasks

import (
	"errors"
)

const (
	DefTaskPoolCount = 200
	TaskSize         = 1024
)

var GlobalTaskPool *TaskPool

var (
	ErrNotFoundCallBack = errors.New("not found callBack function")
)

func init() {
	GlobalTaskPool = newTaskPool(0)
}
