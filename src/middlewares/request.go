package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 初始化请求时间
		t := time.Now()
		// 提取请求体
		requestRawData, err := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewReader(requestRawData))
		if err != nil {
			utils.ContextWarningf(c, "[LoadRequestBodyFailed] %s", err)
		}
		// 执行
		c.Next()
		// 记录请求耗时
		duration := time.Since(t).Milliseconds()
		// 解析请求体
		var requestJsonData map[string]interface{}
		if len(requestRawData) > 0 {
			if err = json.Unmarshal(requestRawData, &requestJsonData); err != nil {
				requestRawData = []byte("-")
			}
		} else {
			requestRawData = []byte("-")
		}
		// 记录请求日志
		utils.ContextInfof(c, "[RequestLog] (%dms) %s %s %s %d", duration, c.Request.Method, c.Request.URL, requestRawData, c.Writer.Status())
	}
}
