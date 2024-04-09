package logging

import (
	"fmt"
	"github.com/jwalton/go-supportscolor"
	"log"
	"time"
)

const UNSET uint8 = 0
const DEBUG uint8 = 10
const INFO uint8 = 20
const WARN uint8 = 30
const ERROR uint8 = 40

type LoggerImpl interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type Options struct {
	Name    string
	Level   uint8
	NoColor bool
}

type Logger struct {
	name    string
	level   uint8
	noColor bool
}

func (logger *Logger) format(level string, message string, args ...interface{}) string {
	timestamp := time.Now().Format(time.DateTime)
	msg := fmt.Sprintf(message, args...)

	if logger.noColor {
		return fmt.Sprintf("%s [%-5s] %s: %s", timestamp, level, logger.name, msg)
	}

	switch level {
	case "DEBUG":
		level = "\x1b[1;47mDEBUG\x1b[0m"
	case "INFO":
		level = "\x1b[1;34mINFO\x1b[0m"
	case "WARN":
		level = "\x1b[1;33mWARN\x1b[0m"
	case "ERROR":
		level = "\x1b[1;31mERROR\x1b[0m"
	}

	return fmt.Sprintf("\x1b[90m%s\x1b[0m %-17s\x1b[36m%s\x1b[0m %s", timestamp, level, logger.name, msg)
}

func (logger *Logger) Debug(msg string, args ...interface{}) {
	if logger.level > 10 {
		return
	}

	formatted := logger.format("DEBUG", msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Info(msg string, args ...interface{}) {
	if logger.level > 20 {
		return
	}

	formatted := logger.format("INFO", msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Warn(msg string, args ...interface{}) {
	if logger.level > 30 {
		return
	}

	formatted := logger.format("WARN", msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Error(msg string, args ...interface{}) {
	if logger.level > 40 {
		return
	}

	formatted := logger.format("ERROR", msg, args...)
	log.Println(formatted)
}

func NewLogger(options Options) LoggerImpl {
	log.SetFlags(0)

	if options.Name == "" {
		options.Name = "__root__"
	}
	logger := &Logger{name: options.Name}

	switch options.Level {
	case 0:
		logger.level = INFO
	case 10:
		logger.level = DEBUG
	case 20:
		logger.level = INFO
	case 30:
		logger.level = WARN
	case 40:
		logger.level = ERROR
	default:
		logger.level = INFO
	}

	if options.NoColor || !supportscolor.Stdout().SupportsColor {
		logger.noColor = true
	} else {
		logger.noColor = false
	}

	return logger
}
