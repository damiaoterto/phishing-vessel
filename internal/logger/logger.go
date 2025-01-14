package logger

import "github.com/fatih/color"

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

func Info(msg string) {
	printLog(INFO, msg)
}

func Error(msg string) {
	printLog(ERROR, msg)
}

func Success(msg string) {
	printLog(SUCCESS, msg)
}

func printLog(level LogLevel, msg string) {
	logColor := logColors[level]
	logColor.Printf("[%s] %s\n", level, msg)
}
