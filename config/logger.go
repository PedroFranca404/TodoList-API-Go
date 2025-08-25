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
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Non-Formated Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}

// Formated Logs
func (l *Logger) Debugf(p string, v ...interface{}) {
	l.debug.Printf(p, v...)
}
func (l *Logger) Infof(p string, v ...interface{}) {
	l.info.Printf(p, v...)
}
func (l *Logger) Warnf(p string, v ...interface{}) {
	l.warning.Printf(p, v...)
}
func (l *Logger) Errf(p string, v ...interface{}) {
	l.err.Printf(p, v...)
}
