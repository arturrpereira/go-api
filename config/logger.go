package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writter io.Writer
}

func NewLogger(p string) *Logger {
	writter := io.Writer(os.Stdout)
	logger := log.New(writter, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writter, "DEBUG: ", logger.Flags()),
		info:    log.New(writter, "INFO: ", logger.Flags()),
		warning: log.New(writter, "WARNING: ", logger.Flags()),
		err:     log.New(writter, "ERROR: ", logger.Flags()),
		writter: writter,
	}
}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
