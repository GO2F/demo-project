package logger

import (
	"rank/config"
    "github.com/lestrrat/go-file-rotatelogs"
    "github.com/pkg/errors"
    "github.com/rifflock/lfshook"
    log "github.com/sirupsen/logrus"
    "path"
    "time"

    "os"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Init 为logrus添加hook
func Init(logPath string) {
    log.SetLevel(log.InfoLevel)
    log.AddHook(newRotateHook(logPath, "rank", 7*24*time.Hour, 24*time.Hour))
}

func newRotateHook(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) *lfshook.LfsHook {
    baseLogPath := path.Join(logPath, logFileName)

    writer, err := rotatelogs.New(
        baseLogPath+".%Y-%m-%d.log",
        rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文
        rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
        rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
    )
    if err != nil {
        log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
    }
    return lfshook.NewHook(lfshook.WriterMap{
        log.DebugLevel: writer, // 为不同级别设置不同的输出目的
        log.InfoLevel:  writer,
        log.WarnLevel:  writer,
        log.ErrorLevel: writer,
        log.FatalLevel: writer,
        log.PanicLevel: writer,
    }, &log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
}


// LogerMiddleware server文件日志中间件
func LogerMiddleware() gin.HandlerFunc {
	// 日志文件
	fileName := path.Join(config.App.LogPathURI, "server_request")
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := log.New()
	//设置日志级别
	logger.SetLevel(log.DebugLevel)
	//设置输出
	logger.Out = src


	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y-%m-%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqURL := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(log.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURL,
		}).Info()

	}
}