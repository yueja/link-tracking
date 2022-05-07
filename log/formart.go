package log

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
	traceutil "github.com/yueja/link-tracking/trace"
)

type logData struct{}

type Server interface {
	setData(args []interface{}) (fields, errFields logrus.Fields)
	setLog(ctx context.Context, level logrus.Level, template string, args []interface{})
}

func newServer() *logData {
	return new(logData)
}

// setData 统一处理 Log 自定义信息
func (l logData) setData(args []interface{}) (fields, errFields logrus.Fields) {
	data := make(map[string]logrus.Fields)
	errFields = make(logrus.Fields)
	for _, v := range args {
		itemFiled, _ := v.(ItemFiled)
		if itemFiled.PKey == itemFiledError {
			errFields[itemFiledError] = itemFiled.Fields[itemFiledError]
			continue
		}
		for fieldKey, fieldValue := range itemFiled.Fields {
			if data[itemFiled.PKey] == nil {
				data[itemFiled.PKey] = make(map[string]interface{})
			}
			data[itemFiled.PKey][fieldKey] = fieldValue
		}
	}
	fields = make(logrus.Fields)
	for key, v := range data {
		fields[key] = v
	}
	return fields, errFields
}

// setLog 日志记录
func (l logData) setLog(ctx context.Context, level logrus.Level, template string, args []interface{}) {
	args = l.setAppName(args)
	fields, errFields := l.setData(args)
	ctxFields := traceutil.NewTraceServer().HandleCtx(ctx)
	entry := log.WithContext(ctx).WithFields(ctxFields).WithFields(fields)
	if len(errFields) != 0 {
		err, ok := errFields[itemFiledError].(error)
		if !ok {
			log.Printf("setLog errFields to error err:%+v", errFields[itemFiledError])
			return
		}
		// 打印堆栈信息
		err1 := errors.Wrap(err, 1)
		fields["stack"] = err1.ErrorStack()
		entry = entry.WithFields(fields).WithError(err)
	}
	switch level {
	case logrus.PanicLevel:
		entry.Panic(template)
	case logrus.FatalLevel:
		entry.Fatal(template)
	case logrus.ErrorLevel:
		entry.Error(template)
	case logrus.WarnLevel:
		entry.Warn(template)
	case logrus.InfoLevel:
		entry.Info(template)
	case logrus.DebugLevel:
		entry.Debug(template)
	case logrus.TraceLevel:
		entry.Trace(template)
	default:
		fmt.Printf("Unknown Log Level:%d,msg:%+v", level, fields)
	}
}

func (l logData) setAppName(args []interface{}) []interface{} {
	return append(args, DAny("app_name", appName))
}
