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

func Info(format string, args ...any) {
	printLog(INFO, format, args...)
}

func Error(format string, args ...any) {
	printLog(ERROR, format, args...)
}

func Success(format string, args ...any) {
	printLog(SUCCESS, format, args...)
}

func printLog(level LogLevel, format string, args ...any) {
	logColor := logColors[level]
	logColor.Printf("[%s] %v\n", level, fmt.Sprintf(format, args...))
}
