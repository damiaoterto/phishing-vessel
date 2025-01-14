package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	SUCCESS
	DEFAULT
)

func (l LogLevel) String() string {
	return [...]string{"INFO", "ERROR", "SUCCESS", "DEFAULT"}[l]
}

var logColors = map[LogLevel]*color.Color{
	INFO:    color.New(color.FgCyan),
	ERROR:   color.New(color.FgRed),
	SUCCESS: color.New(color.FgGreen),
	DEFAULT: color.New(color.FgWhite),
}

func Infof(format string, args ...any) {
	printLogf(INFO, format, args...)
}

func Info(msg string) {
	printLog(INFO, msg)
}

func Errorf(format string, args ...any) {
	printLogf(ERROR, format, args...)
}

func Error(msg string) {
	printLog(ERROR, msg)
}

func Successf(format string, args ...any) {
	printLogf(SUCCESS, format, args...)
}

func Success(msg string) {
	printLog(SUCCESS, msg)
}

func printLogf(level LogLevel, format string, args ...any) {
	logColor := logColors[level]
	logColor.Printf("[%s] %v\n", level, fmt.Sprintf(format, args...))
}

func printLog(level LogLevel, msg string) {
	logColor := logColors[level]
	logColor.Printf("[%s] %s\n", level, msg)
}
