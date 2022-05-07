package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yueja/link-tracking/log"
	"io/ioutil"
	"time"
)

func VisitLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := ctx.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		body := UnmarshalJsonBody(data)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		sessionId := ctx.Request.Header.Get("Authorization")
		start := time.Now()
		log.InfoW(
			ctx, "VisitRequest",
			log.DAny("method", ctx.Request.Method),
			log.DAny("uri", ctx.FullPath()),
			log.DAny("ip", ctx.ClientIP()),
			log.DAny("sid", sessionId),
			log.DAny("body", body),
			log.DAny("full_uri", ctx.Request.RequestURI),
			log.DAny("header", ctx.Request.Header),
		)
		respWriter := NewResponseWriter(ctx)
		ctx.Writer = respWriter
		ctx.Next()
		latency := time.Now().Sub(start).Nanoseconds()
		resBody := respWriter.body.Bytes()
		log.InfoW(
			ctx, "VisitResponse",
			log.DAny("uri", ctx.Request.RequestURI),
			log.DAny("status", ctx.Writer.Status()),
			log.DAny("size", ctx.Writer.Size()),
			log.DAny("latency", latency),
			log.DAny("failed", ctx.Errors.Errors()),
			log.DAny("resp_body", UnmarshalJsonBody(resBody)),
		)
	}
}

// UnmarshalJsonBody 将body Unmarshal 为 Json
func UnmarshalJsonBody(body []byte) (jsonBody map[string]interface{}) {
	jsonBody = make(map[string]interface{})
	err := json.Unmarshal(body, &jsonBody)
	if err != nil { // 当不为 JsonBody 时
		jsonBody["_text"] = string(body)
	}
	return jsonBody
}

func NewResponseWriter(ctx *gin.Context) *ResponseWriter {
	return &ResponseWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: ctx.Writer,
	}
}

// ResponseWriter 记录每个请求的日志
type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
