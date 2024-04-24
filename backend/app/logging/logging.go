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
const CRITICAL uint8 = 50

type LoggerImpl interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Critical(msg string, args ...interface{})
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

func (logger *Logger) format(level uint8, message string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05.0000")
	msg := fmt.Sprintf(message, args...)

	var levelS string
	var levelN string

	switch level {
	case DEBUG:
		levelS = "\x1b[1;47mDEBUG\x1b[0m"
		levelN = "DEBUG"
	case INFO:
		levelS = "\x1b[1;34mINFO\x1b[0m"
		levelN = "INFO"
	case WARN:
		levelS = "\x1b[1;33mWARN\x1b[0m"
		levelN = "WARN"
	case ERROR:
		levelS = "\x1b[1;31mERROR\x1b[0m"
		levelN = "ERROR"
	case CRITICAL:
		levelN = "CRITICAL"
	default:
		levelS = "\x1b[1;47mUNSET\x1b[0m"
		levelN = "UNSET"
	}

	if logger.noColor {
		return fmt.Sprintf("%s [%-8s] %s: %s", timestamp, levelN, logger.name, msg)
	}

	if level == CRITICAL {
		return fmt.Sprintf("\x1b[1;31m%s CRITICAL\x1b[0m \x1b[36m%s\x1b[0m %s", timestamp, logger.name, msg)
	}

	return fmt.Sprintf("\x1b[90m%s\x1b[0m %-20s\x1b[36m%s\x1b[0m %s", timestamp, levelS, logger.name, msg)
}

func (logger *Logger) Debug(msg string, args ...interface{}) {
	if logger.level > 10 {
		return
	}

	formatted := logger.format(DEBUG, msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Info(msg string, args ...interface{}) {
	if logger.level > 20 {
		return
	}

	formatted := logger.format(INFO, msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Warn(msg string, args ...interface{}) {
	if logger.level > 30 {
		return
	}

	formatted := logger.format(WARN, msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Error(msg string, args ...interface{}) {
	if logger.level > 40 {
		return
	}

	formatted := logger.format(ERROR, msg, args...)
	log.Println(formatted)
}

func (logger *Logger) Critical(msg string, args ...interface{}) {
	if logger.level > 50 {
		return
	}

	formatted := logger.format(CRITICAL, msg, args...)
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
	case 50:
		logger.level = CRITICAL
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
