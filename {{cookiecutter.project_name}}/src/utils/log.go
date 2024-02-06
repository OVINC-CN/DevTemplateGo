package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"os"
	"time"
)

var (
	Logger   *logrus.Entry
	DbLogger *DBLoggerType
)

func InitLogger(debug bool, level logrus.Level) *logrus.Entry {
	rawLogger := logrus.New()
	rawLogger.SetLevel(level)
	if debug {
		rawLogger.SetFormatter(&logrus.TextFormatter{
			DisableQuote:     false,
			DisableTimestamp: false,
			TimestampFormat:  time.RFC3339,
		})
		rawLogger.SetOutput(os.Stdout)
	} else {
		rawLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339,
			DisableTimestamp:  false,
			DisableHTMLEscape: true,
			PrettyPrint:       debug,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		})
		file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		rawLogger.SetOutput(file)
	}
	logger := rawLogger.WithFields(logrus.Fields{"request_id": ""})
	logger.Infof("[InitLoggerSuccess] %T", Logger)
	return logger
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

type DBLoggerType struct {
	Logger                              *logrus.Entry
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	slowThreshold                       time.Duration
}

func (l *DBLoggerType) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	return &newLogger
}

func (l *DBLoggerType) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *DBLoggerType) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Warningf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *DBLoggerType) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
}

func (l *DBLoggerType) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case elapsed > l.slowThreshold && l.slowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			l.Logger.Warningf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.Warningf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.Logger.Level == logrus.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Logger.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func InitDBLogger(debug bool, level logrus.Level, slowThreshold time.Duration) *DBLoggerType {
	logger := InitLogger(debug, level)
	dbLogger := &DBLoggerType{
		Logger:        logger,
		slowThreshold: slowThreshold,
		infoStr:       "[Database] %s",
		warnStr:       "[Database] %s",
		errStr:        "[Database] %s",
		traceStr:      "[DatabaseTrace] %s\n[%.3fms] [rows:%v] %s",
		traceWarnStr:  "[DatabaseTrace] %s %s\n[%.3fms] [rows:%v] %s",
		traceErrStr:   "[DatabaseTrace] %s %s\n[%.3fms] [rows:%v] %s",
	}
	return dbLogger
}
