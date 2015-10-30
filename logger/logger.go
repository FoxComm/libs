package logger

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DEBUG = iota + 1
	INFO
	WARN
	ERROR
	PANIC
	FATAL
)

var (
	Level   int
	mapping = map[int]logrus.Level{
		DEBUG: logrus.DebugLevel,
		INFO:  logrus.InfoLevel,
		WARN:  logrus.WarnLevel,
		ERROR: logrus.ErrorLevel,
		PANIC: logrus.PanicLevel,
		FATAL: logrus.FatalLevel,
	}
)

type Map map[string]interface{}

func init() {
	Level = level()
	logrus.SetLevel(mapping[Level])
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func addContext(entry *logrus.Entry) *logrus.Entry {
	for k, v := range context() {
		if entry.Data[k] == nil {
			entry.Data[k] = v
		}
	}
	return entry
}

func errorMessage(obj interface{}) (msg string) {
	switch obj.(type) {
	case string:
		msg = obj.(string)
	case error:
		msg = obj.(error).Error()
	}
	return
}

func level() (level int) {
	level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		level = 1
	}

	return
}

func withArgs(msg string, args interface{}) (message string, entry *logrus.Entry) {
	var fields Map
	data := args.([]interface{})

	if len(data) == 0 {
		message = msg
		entry = withFields(Map{})
		return
	}

	switch data[0].(type) {
	case Map:
		message = msg
		fields = data[0].(Map)
	default:
		message = fmt.Sprintf(msg, data...)
		fields = Map{}
	}

	entry = withFields(fields)

	return
}

func withFields(fields map[string]interface{}) *logrus.Entry {
	return logrus.WithFields(fields)
}

func Debug(msg string, args ...interface{}) {
	message, entry := withArgs(msg, args)
	entry.Debug(message)
}

func Error(msg interface{}, args ...interface{}) {
	message, entry := withArgs(errorMessage(msg), args)
	addContext(entry).Error(message)
}

func Fatal(msg string, args ...interface{}) {
	message, entry := withArgs(errorMessage(msg), args)
	addContext(entry).Fatal(message)
}

func Info(msg string, args ...interface{}) {
	message, entry := withArgs(msg, args)
	entry.Info(message)
}

func Panic(msg interface{}, args ...interface{}) {
	message, entry := withArgs(errorMessage(msg), args)
	addContext(entry).Panic(message)
}

func Warn(msg string, args ...interface{}) {
	message, entry := withArgs(msg, args)
	addContext(entry).Warn(message)
}

func LogWithContext(msg string, severity int, ctx LogContext, args ...interface{}) {
	message, entry := withArgs(msg, args)

	for k, v := range ctx.Context() {
		if entry.Data[k] == nil {
			entry.Data[k] = v
		}
	}

	// entry.log is private

	switch severity {
	case DEBUG:
		entry.Debug(message)
	case INFO:
		entry.Info(message)
	case WARN:
		entry.Warn(message)
	case ERROR:
		entry.Error(message)
	case PANIC:
		entry.Panic(message)
	case FATAL:
		entry.Fatal(message)
	default:
		entry.Data["severityUnknown"] = true
		entry.Error(message)
	}
}
