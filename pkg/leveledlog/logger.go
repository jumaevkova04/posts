package leveledlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type Level int8

const (
	LevelAll Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	out      io.Writer
	minLevel Level
	useJSON  bool
	mu       sync.Mutex
}

func NewLogger(out io.Writer, minLevel Level) *Logger {
	var logger = &Logger{
		out:      out,
		minLevel: minLevel,
	}
	l := strings.Repeat("-", 33)
	logger.Info(fmt.Sprintf("%s%s%s", l, "Starting Logging", l))
	return logger
}

func (l *Logger) Info(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.print(LevelInfo, message)
}

func (l *Logger) Warning(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.print(LevelWarning, message)
}

func (l *Logger) Error(err error) {
	l.print(LevelError, err.Error())
}

func (l *Logger) Fatal(err error) {
	l.print(LevelFatal, err.Error())
	os.Exit(1)
}

func (l *Logger) print(level Level, message string) {
	if level < l.minLevel {
		return
	}

	var line string

	if l.useJSON {
		line = jsonLine(level, message)
	} else {
		line = textLine(level, message)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	_, _ = fmt.Fprintln(l.out, line)
}

func textLine(level Level, message string) string {
	t := time.Now()
	line := fmt.Sprintf("%v %v %q", t.Format("2006/01/02 15:04:05"), level, message)

	if level >= LevelError {
		line += fmt.Sprintf("\n%s", string(debug.Stack()))
	}

	return line
}

func jsonLine(level Level, message string) string {
	aux := struct {
		Level   string `json:"level"`
		Time    string `json:"time"`
		Message string `json:"message"`
		Trace   string `json:"trace,omitempty"`
	}{
		Level:   level.String(),
		Time:    time.Now().UTC().Format(time.RFC3339),
		Message: message,
	}

	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	var line []byte

	line, err := json.Marshal(aux)
	if err != nil {
		return fmt.Sprintf("%s: unable to marshal log message: %s", LevelError.String(), err.Error())
	}

	return string(line)
}
