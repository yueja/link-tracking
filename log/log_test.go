package log

import (
	"context"
	"errors"
	"testing"
)

func Test_Info(t *testing.T) {
	type Temp struct {
		Name string `json:"name"`
		Tel  string `json:"tel"`
	}
	temp := Temp{Name: "yueja", Tel: "12345"}
	InfoW(context.Background(), "我是一条文案哦：", DAny("test", 111), DAny("haha", temp))
}

func Test_Err(t *testing.T) {
	ctx := context.Background()
	type Temp struct {
		Name string `json:"name"`
		Tel  string `json:"tel"`
	}
	temp := Temp{Name: "yueja", Tel: "12345"}
	err := errors.New("哈哈哈哈哈哈哈哈哈")
	ErrorW(ctx, "我是一条文案哦：",
		DAny("test", 111),
		DAny("haha", temp),
		DErr(err),
	)
}
