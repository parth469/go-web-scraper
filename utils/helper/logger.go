package helper

import (
	"fmt"
	"github.com/phuslu/log"
	"io"
	"strings"
)

type Logger struct {
	Logger log.Logger
}

func (l *Logger) Info(fmt string, data ...any) { l.Logger.Info().Msgf(fmt, data...) }

func (l *Logger) Error(format string, err error, args ...any) {
	if err != nil {
		l.Logger.Error().Msgf(format+": %v", append(args, err)...)
	} else {
		l.Logger.Error().Msgf(format, args...)
	}
}

func (l *Logger) Debug(format string, args ...any) {
	l.Logger.Debug().Msgf(format, args...)
}

func (l *Logger) Warn(format string, args ...any) {
	l.Logger.Warn().Msgf(format, args...)
}

func (l *Logger) Fatal(format string, error error, args ...any) {
	if error != nil {
		l.Logger.Fatal().Msgf(format+": %v", append(args, error)...)
	} else {
		l.Logger.Fatal().Msgf(format, args...)
	}
}

func Formatter(w io.Writer, a *log.FormatterArgs) (int, error) {
	levelColor := map[string]string{
		"debug": "\033[36m",   // Cyan
		"info":  "\033[32m",   // Green
		"warn":  "\033[33m",   // Yellow
		"error": "\033[31m",   // Red
		"fatal": "\033[31;1m", // Dark Red (Bold Red)
	}

	reset := "\033[0m"
	color := levelColor[strings.ToLower(a.Level)]

	return fmt.Fprintf(
		w,
		"%s[%s] | %-5s | %s%s\n",
		color,
		a.Time,
		strings.ToUpper(a.Level),
		a.Message,
		reset,
	)
}

var Log = &Logger{log.Logger{
	Level:      log.InfoLevel,
	TimeFormat: "02-01 15:04:05",
	Caller:     1,
	Writer: &log.ConsoleWriter{
		Formatter: Formatter,
	},
}}
