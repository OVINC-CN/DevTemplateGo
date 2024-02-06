package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	Logger *logrus.Entry
)

func InitLogger(debug bool, level logrus.Level) {
	rawLogger := logrus.New()
	rawLogger.SetLevel(level)
	rawLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		PrettyPrint:       debug,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})
	if debug {
		rawLogger.SetOutput(os.Stdout)
	} else {
		file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		rawLogger.SetOutput(file)
	}
	Logger = rawLogger.WithFields(logrus.Fields{"request_id": ""})
	Logger.Infof("[InitLoggerSuccess] %T", Logger)
}

func getContextLogger(c *gin.Context) (logger *logrus.Entry, exists bool) {
	val, ok := c.Get("logEntry")
	if !ok {
		Logger.Warnf("[RequestLoggerNotSet] %s", val)
		return
	}
	logger, exists = val.(*logrus.Entry)
	return
}

func ContextInfof(c *gin.Context, format string, args ...interface{}) {
	logger, _ := getContextLogger(c)
	logger.Infof(format, args...)
}

func ContextErrorf(c *gin.Context, format string, args ...interface{}) {
	logger, _ := getContextLogger(c)
	logger.Errorf(format, args...)
}

func ContextWarningf(c *gin.Context, format string, args ...interface{}) {
	logger, _ := getContextLogger(c)
	logger.Warningf(format, args...)
}
