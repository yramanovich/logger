package logger

import (
	"io"
	"os"
	"time"
)

// Logger is a common logging contract.
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
}

type abstractLogger struct {
	ef       Formatter
	w        io.Writer
	logLevel LogLevel
	errc     chan<- error
}

func (a abstractLogger) Info(args ...interface{}) {
	a.logEvent(LevelInfo, args...)
}

func (a abstractLogger) Warn(args ...interface{}) {
	a.logEvent(LevelWarn, args...)
}

func (a abstractLogger) Error(args ...interface{}) {
	a.logEvent(LevelError, args...)
}

func (a abstractLogger) Fatal(args ...interface{}) {
	a.logEvent(LevelFatal, args...)
}

func (a abstractLogger) Debug(args ...interface{}) {
	a.logEvent(LevelDebug, args...)
}

func (a abstractLogger) Trace(args ...interface{}) {
	a.logEvent(LevelTrace, args...)
}

func (a abstractLogger) logEvent(evLevel LogLevel, args ...interface{}) {
	if !isLogAllowed(a.logLevel, evLevel) {
		return
	}
	msg, err := a.ef.message(evLevel, time.Now(), args...)
	if err != nil {
		go a.errNotify(err)
		return
	}
	if _, err = io.WriteString(a.w, msg); err != nil {
		go a.errNotify(err)
	}
}

func (a abstractLogger) errNotify(err error) {
	if err != nil && a.errc != nil {
		a.errc <- err
	}
}

// New creates new Logger instance.
// If no options provided uses default formatter, INFO log level and writes logs to os.Stdout.
func New(opts ...func(al *abstractLogger)) Logger {
	al := abstractLogger{
		ef:       defaultFormatter{},
		w:        os.Stdout,
		logLevel: LevelInfo,
		errc:     nil,
	}
	for _, opt := range opts {
		opt(&al)
	}
	return al
}

// SetLogLevel sets logging level.
func SetLogLevel(level LogLevel) func(al *abstractLogger) {
	return func(al *abstractLogger) {
		al.logLevel = level
	}
}

// SetFormatter sets formatter implementation.
func SetFormatter(formatter Formatter) func(al *abstractLogger) {
	return func(al *abstractLogger) {
		al.ef = formatter
	}
}

// SetErrChannel sets error channel, so logger can notify about failures.
func SetErrChannel(errc chan<- error) func(al *abstractLogger) {
	return func(al *abstractLogger) {
		al.errc = errc
	}
}

// SetWriter sets output writer where log messages go.
func SetWriter(w io.Writer) func(al *abstractLogger) {
	return func(al *abstractLogger) {
		al.w = w
	}
}
