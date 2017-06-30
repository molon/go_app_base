package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// LogLevel log level
type LogLevel int

const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
	FATAL = LogLevel(5)
)

// Logger is the server logger
type Logger struct {
	logger     *log.Logger
	logLevel   LogLevel
	debugLabel string
	infoLabel  string
	warnLabel  string
	errorLabel string
	fatalLabel string
}

// Generate the pid prefix string
func pidPrefix() string {
	return fmt.Sprintf("[%d] ", os.Getpid())
}

func setLabelFormats(l *Logger, colors bool) {
	if colors {
		colorFormat := "[\x1b[%dm%s\x1b[0m] "
		l.debugLabel = fmt.Sprintf(colorFormat, 36, "DBG")
		l.infoLabel = fmt.Sprintf(colorFormat, 32, "INF")
		l.warnLabel = fmt.Sprintf(colorFormat, 33, "WRN")
		l.errorLabel = fmt.Sprintf(colorFormat, 31, "ERR")
		l.fatalLabel = fmt.Sprintf(colorFormat, 35, "FTL")
	} else {
		l.debugLabel = "[DBG] "
		l.infoLabel = "[INF] "
		l.warnLabel = "[WRN] "
		l.errorLabel = "[ERR] "
		l.fatalLabel = "[FTL] "
	}
}

// NewStdLogger creates a logger with output directed to Stderr
func NewStdLogger(time, colors, pid bool, logLevel LogLevel) *Logger {
	flags := 0
	if time {
		flags = log.LstdFlags | log.Lmicroseconds
	}

	pre := ""
	if pid {
		pre = pidPrefix()
	}

	l := &Logger{
		logger:   log.New(os.Stderr, pre, flags),
		logLevel: logLevel,
	}

	setLabelFormats(l, colors)

	return l
}

// ParseLogLevel parse a levvel str to LogLevel
func ParseLogLevel(levelstr string) (LogLevel, error) {
	lvl := INFO

	switch strings.ToLower(levelstr) {
	case "debug":
		lvl = DEBUG
	case "info":
		lvl = INFO
	case "warn":
		lvl = WARN
	case "error":
		lvl = ERROR
	case "fatal":
		lvl = FATAL
	default:
		return lvl, fmt.Errorf("invalid log-level '%s'", levelstr)
	}

	return lvl, nil
}

// Logf printf log with level and format
func (l *Logger) Logf(msgLevel LogLevel, format string, v ...interface{}) {
	if l.logLevel > msgLevel {
		return
	}

	label := ""

	switch msgLevel {
	case DEBUG:
		label = l.debugLabel
	case INFO:
		label = l.infoLabel
	case WARN:
		label = l.warnLabel
	case ERROR:
		label = l.errorLabel
	case FATAL:
		label = l.fatalLabel
	}

	l.logger.Printf(label+format, v...)
}
