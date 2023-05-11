package log

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	serviceName    = "todolist-http"
	developmentEnv = "development"
	stagingEnv     = "staging"
	productionEnv  = "production"
	appEnv         = "APPENV"
	logPath        = "/var/log/app/"
)

var (
	ErrLog    = logrus.New()
	AccessLog = logrus.New()
	DebugLog  = logrus.New()
)

func init() {
	initErrorLogFile()
	initAccessLogFile()
	initDebugLogFile()
}

func GetErrorLogFile() string {
	var errLogPath, env string

	env = os.Getenv(appEnv)
	if env == "" {
		env = developmentEnv
	}

	switch env {
	case developmentEnv:
		errLogPath = "./" + serviceName + "." + developmentEnv + ".error.log"
	case stagingEnv, productionEnv:
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			ErrLog.Errorln("failed initializing log folder")
		}
		errLogPath = fmt.Sprintf("%s%s.%s.error.log", logPath, serviceName, env)
	}

	return errLogPath
}

func GetAccessLogFile() string {
	var accLogPath, env string

	env = os.Getenv(appEnv)
	if env == "" {
		env = developmentEnv
	}

	switch env {
	case developmentEnv:
		accLogPath = "./" + serviceName + "." + developmentEnv + ".access.log"
	case stagingEnv, productionEnv:
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			ErrLog.Errorln("failed initializing log folder")
		}
		accLogPath = fmt.Sprintf("%s%s.%s.access.log", logPath, serviceName, env)
	}

	return accLogPath
}

func initErrorLogFile() {
	logPath := GetErrorLogFile()

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		ErrLog.Out = file
	} else {
		ErrLog.Out = os.Stdout
		ErrLog.Error("failed initiating error log file, will use default stderr..")
	}
}

func initAccessLogFile() {
	logPath := GetAccessLogFile()

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		AccessLog.Out = file
	} else {
		AccessLog.Out = os.Stdout
		AccessLog.Error("failed initiating access log file, will use default stderr..")
	}
}

func initDebugLogFile() {
	logFormatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}

	DebugLog.SetLevel(logrus.DebugLevel)
	DebugLog.SetFormatter(logFormatter)
}

func FieldsCtx(ctx context.Context, fields logrus.Fields) logrus.Fields {
	rawRequestID := ctx.Value("request_id")
	requestID, ok := rawRequestID.(string)
	if ok {
		fields["request_id"] = requestID
	}

	return fields
}
