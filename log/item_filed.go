package log

import (
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

const (
	itemFiledData  = "data"
	itemFiledError = "error"
)

type ItemFiled struct {
	PKey          string // 父级 field key
	logrus.Fields        // 子级 Field
}

// DAny 用于 `data` 中的 ItemFiled 打印
func DAny(key string, data interface{}) ItemFiled {
	return ItemAny(itemFiledData, key, data)
}

// DErr 用于 `error` 中的 ItemFiled 打印
func DErr(data interface{}) ItemFiled {
	return ItemAny(itemFiledError, itemFiledError, data)
}

// WithStack 日志堆栈信息
func WithStack(err error) error {
	err = errors.Wrap(err, 1)
	return err
}

// ItemAny 创建一个ItemFiled
// 用于生成任意类型的指定 pKey 的 ItemFiled
func ItemAny(pKey string, key string, val interface{}) ItemFiled {
	return ItemFiled{
		PKey: pKey,
		Fields: map[string]interface{}{
			key: val,
		},
	}
}
