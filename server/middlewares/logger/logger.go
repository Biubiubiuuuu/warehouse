package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/fileHelper"
	"github.com/gin-gonic/gin"
)

// 文件路径，以天为目录
var path = configHelper.LogDir + time.Now().Format("20060102") + "/"
var logChannel = make(chan string, 100)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	go handleLogChannel()
	return func(c *gin.Context) {
		// 记录请求的url、method、header、postData、responseData
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// 响应接收时间
		stratTime := time.Now().String()
		request := c.Request
		var bodyBytes []byte
		var postData string
		if request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(request.Body)
		}
		var postDataJson bytes.Buffer
		if err := json.Indent(&postDataJson, []byte(bodyBytes), "", "\t"); err != nil {
			postData = string(bodyBytes)
		} else {
			postData = postDataJson.String()
		}
		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// url
		url := request.Host + request.URL.String()
		// method
		method := request.Method
		// Content-Type
		contentTypeArr := request.Header["Content-Type"]
		var contentType string
		if len(contentTypeArr) > 0 {
			contentType = contentTypeArr[0]
		}
		// Authorization
		authorizationArr := request.Header["Authorization"]
		var authorization string
		if len(authorizationArr) > 0 {
			authorization = authorizationArr[0]
		}
		c.Next()
		// responseData
		responseBody := bodyLogWriter.body.String()
		// 响应返回时间
		endTime := time.Now().String()
		log := fmt.Sprintf("开始时间：%s\r\n请求URL：%s\r\nmethod：%s\r\n%s：%s\r\n%s：%s\r\n请求IP：%s\r\npostData：\r\n%s\r\nresponseBody：\r\n%s\r\n结束时间：%s\r\n", stratTime, url, method, "Content-Type", contentType, "Authorization", authorization, c.ClientIP(), postData, responseBody, endTime)
		logChannel <- log
	}
}

func handleLogChannel() {
	if !fileHelper.IsExist(path) {
		fileHelper.CreateDir(path)
	}
	fileName := path + time.Now().Format("20060102") + ".log"
	if f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {

	} else {
		for log := range logChannel {
			_, _ = f.WriteString("========================================\r\n")
			_, _ = f.WriteString(log)
			_, _ = f.WriteString("========================================\r\n")
		}
	}
}
