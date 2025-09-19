package logger

import (
	"log"
	"os"
	"time"
)

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger provides structured logging functionality
type Logger struct {
	level  LogLevel
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", message, fields...)
	}
}

// Info logs an info message
func (l *Logger) Info(message string, fields ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", message, fields...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(message string, fields ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", message, fields...)
	}
}

// Error logs an error message
func (l *Logger) Error(message string, fields ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", message, fields...)
	}
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(message string, fields ...interface{}) {
	l.log("FATAL", message, fields...)
	os.Exit(1)
}

// log formats and logs a message
func (l *Logger) log(level, message string, fields ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	if len(fields) > 0 {
		l.logger.Printf("[%s] %s: %s %v", timestamp, level, message, fields)
	} else {
		l.logger.Printf("[%s] %s: %s", timestamp, level, message)
	}
}

// Global logger instance
var defaultLogger = NewLogger(INFO)

// SetGlobalLevel sets the global logging level
func SetGlobalLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

// Debug logs a debug message using the global logger
func Debug(message string, fields ...interface{}) {
	defaultLogger.Debug(message, fields...)
}

// Info logs an info message using the global logger
func Info(message string, fields ...interface{}) {
	defaultLogger.Info(message, fields...)
}

// Warn logs a warning message using the global logger
func Warn(message string, fields ...interface{}) {
	defaultLogger.Warn(message, fields...)
}

// Error logs an error message using the global logger
func Error(message string, fields ...interface{}) {
	defaultLogger.Error(message, fields...)
}

// Fatal logs a fatal message using the global logger and exits
func Fatal(message string, fields ...interface{}) {
	defaultLogger.Fatal(message, fields...)
}
