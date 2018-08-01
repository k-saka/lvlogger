// Package lvlogger provides level base logger.
// lvlogger is implemented as simple log.Logger wrapper.
package lvlogger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Log levels
const (
	Debug int8 = iota
	Info
	Warn
	Error
	Fatal
)

// UnmarshalLogLevel converts log level string to int8 log level
func UnmarshalLogLevel(levelStr string) (int8, error) {
	s := strings.ToLower(levelStr)
	switch s {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	default:
		return 0, errors.New("Invalid level string:" + levelStr)
	}
}

// LvLogger is level based logger.
// It wrapps go standard log.Logger with log level functions.
type LvLogger struct {
	inner *log.Logger
	level int8
}

// NewLvLogger creates a new LvLogger.
// The level argument defines effective log level.
// The other argments pass to log.Logger.New.
func NewLvLogger(out io.Writer, prefix string, flag int, level int8) *LvLogger {
	inner := log.New(out, prefix, flag)
	return &LvLogger{inner: inner, level: level}
}

// Debugf logs a message if l.level is Debug.
// Arguments are handled in the manner of fmt.Printf.
func (l *LvLogger) Debugf(format string, v ...interface{}) {
	if l.level > Debug {
		return
	}

	l.inner.Output(2, fmt.Sprintf(format, v...))
}

// Debug logs a message if l.level is Debug.
// Arguments are handled in the manner of fmt.Print.
func (l *LvLogger) Debug(v ...interface{}) {
	if l.level > Debug {
		return
	}

	l.inner.Output(2, fmt.Sprint(v...))
}

// Debugln logs a message if l.level is Debug.
// Arguments are handled in the manner of fmt.Println.
func (l *LvLogger) Debugln(v ...interface{}) {
	if l.level > Debug {
		return
	}
	l.inner.Output(2, fmt.Sprintln(v...))
}

// Infof logs a message if l.level is Info or less.
// Arguments are handled in the manner of fmt.Printf.
func (l *LvLogger) Infof(format string, v ...interface{}) {
	if l.level > Info {
		return
	}

	l.inner.Output(2, fmt.Sprintf(format, v...))
}

// Info logs a message if l.level is less Info or less.
// Arguments are handled in the manner of fmt.Print.
func (l *LvLogger) Info(v ...interface{}) {
	if l.level > Info {
		return
	}

	l.inner.Output(2, fmt.Sprint(v...))
}

// Infoln logs a message if l.level is Info or less.
// Arguments are handled in the manner of fmt.Println.
func (l *LvLogger) Infoln(v ...interface{}) {
	if l.level > Info {
		return
	}
	l.inner.Output(2, fmt.Sprintln(v...))
}

// Warnf logs a message if l.level is Warn or less.
// Arguments are handled in the manner of fmt.Printf.
func (l *LvLogger) Warnf(format string, v ...interface{}) {
	if l.level > Warn {
		return
	}

	l.inner.Output(2, fmt.Sprintf(format, v...))
}

// Warn logs a message if l.level is less Warn or less.
// Arguments are handled in the manner of fmt.Print.
func (l *LvLogger) Warn(v ...interface{}) {
	if l.level > Warn {
		return
	}

	l.inner.Output(2, fmt.Sprint(v...))
}

// Warnln logs a message if l.level is Warn or less.
// Arguments are handled in the manner of fmt.Println.
func (l *LvLogger) Warnln(v ...interface{}) {
	if l.level > Warn {
		return
	}
	l.inner.Output(2, fmt.Sprintln(v...))
}

// Errorf logs a message if l.level is Error or less.
// Arguments are handled in the manner of fmt.Printf.
func (l *LvLogger) Errorf(format string, v ...interface{}) {
	if l.level > Error {
		return
	}

	l.inner.Output(2, fmt.Sprintf(format, v...))
}

// Error logs a message if l.level is less Error or less.
// Arguments are handled in the manner of fmt.Print.
func (l *LvLogger) Error(v ...interface{}) {
	if l.level > Error {
		return
	}

	l.inner.Output(2, fmt.Sprint(v...))
}

// Errorln logs a message if l.level is Error or less.
// Arguments are handled in the manner of fmt.Println.
func (l *LvLogger) Errorln(v ...interface{}) {
	if l.level > Error {
		return
	}
	l.inner.Output(2, fmt.Sprintln(v...))
}

// Fatalf is quivalent to log.Logger.Fatalf.
func (l *LvLogger) Fatalf(format string, v ...interface{}) {
	l.inner.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatal is quivalent to log.Logger.Fatal.
func (l *LvLogger) Fatal(v ...interface{}) {
	l.inner.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalln is quivalent to log.Logger.Fatalln.
func (l *LvLogger) Fatalln(v ...interface{}) {
	l.inner.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
