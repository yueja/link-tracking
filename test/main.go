package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yueja/link-tracking/log"
	midLog "github.com/yueja/link-tracking/middleware/log"
	"github.com/yueja/link-tracking/middleware/trace"
	"net/http"
	"time"
)

func main() {
	log.Init(log.LoggerConfig{AppName: "crm"})
	r := gin.Default()
	r.Use(trace.WithTrace())
	r.Use(midLog.VisitLog())
	r.GET("/index", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		var student struct {
			Name string `json:"name"`
			Tel  string `json:"tel"`
		}
		student.Tel = "18000000000"
		student.Name = "李四"
		log.InfoW(ctx, "测试info日志", log.DAny("name", "张三"), log.DAny("student", student))
		log.ErrorW(ctx, "测试自定义错误", log.DAny("年龄", 20), log.DErr(errors.New("测试自定义错误")))
		if err := test1(); err != nil {
			log.ErrorW(ctx, "测试堆栈信息", log.DAny("年龄", 20), log.DErr(err))
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}

func test1() (err error) {
	return test2()
}

func test2() (err error) {
	err = errors.New("哈哈哈哈哈test")
	return log.WithStack(err)
}
