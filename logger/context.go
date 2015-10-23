package logger

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"runtime"
	"strings"
)

type LogContext struct {
	FuncRegexp *regexp.Regexp
	FileRegexp *regexp.Regexp
	PackRegexp *regexp.Regexp
	SkipStack  int
}

var (
	defaultCtx = LogContext{
		FuncRegexp: regexp.MustCompile(`(?i)FoxComm\/(FoxComm|libs|core_services)\/(\w+.+)\.(.+)$`),
		FileRegexp: regexp.MustCompile(`(?i)FoxComm\/(FoxComm|libs|core_services)\/(\w+.+\.\w+)$`),
		PackRegexp: regexp.MustCompile(`(^\w+)\.(.+)`),
		SkipStack:  3,
	}
)

func context() Map {
	return defaultCtx.Context()
}

func (ctx LogContext) Context() Map {
	var name string
	var funcResult, fileResult, packResult []string

	pc, file, line, ok := runtime.Caller(ctx.SkipStack)

	context := Map{}

	if ok == false {
		logrus.Warn("Unable to gather runtime context", logrus.Fields{
			"function": "context",
			"package":  "social_analytics/logger",
		})
		return context
	}

	name = runtime.FuncForPC(pc).Name()

	fileResult = ctx.FileRegexp.FindStringSubmatch(file)

	context["file"] = fileResult[1]
	context["line"] = line

	if strings.Contains(name, "main") {
		packResult = ctx.PackRegexp.FindStringSubmatch(name)
		context["package"] = packResult[1]
		context["function"] = packResult[2]
	} else {
		funcResult = ctx.FuncRegexp.FindStringSubmatch(name)
		context["package"] = funcResult[1]
		context["function"] = funcResult[2]
	}

	return context
}
