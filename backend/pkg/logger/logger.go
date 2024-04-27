package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Info(message string)
	Warn(context string, message string)
	Error(context string, err error)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l logger) Info(message string) {
	fmt.Printf(
		"%s [INFO] [Function: %s] %s\n",
		time.Now().Format("01/02/2006 03:04pm"), caller(), message,
	)
}

func (l logger) Warn(context string, message string) {
	fmt.Printf(
		"%s [WARNING] [Function: %s (Context: %s)] %s\n",
		time.Now().Format("01/02/2006 03:04pm"), caller(), context, message,
	)
}

func (l logger) Error(context string, err error) {
	fmt.Printf(
		"%s [ERROR] [Function: %s (Context: %s)] %s\n",
		time.Now().Format("01/02/2006 03:04pm"), caller(), context, err,
	)
}

func caller() string {
	count, _, _, success := runtime.Caller(2)
	if !success {
		return "unknown"
	}
	name := runtime.FuncForPC(count).Name()
	return name[strings.LastIndex(name, "/")+1:]
}
