// app_logger package contains definition and logic for custom minimal logger.
package app_logger

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/aantoschuk/feed/internal/apperr"
)

var logLevel = new(slog.LevelVar)

// AppLogger object. Based on the slog with additional parameter
// to handle amount of visible  information in the console.
type AppLogger struct {
	// slog
	logger *slog.Logger
	// controls amount of visible information basic | info | debug
	mode string
}

func loggerHandler(w io.Writer) slog.Handler {
	return slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{} // remove timestamp
			}
			return a
		},
	})
}

// NewAppLogger function allocates AppLogger.
// Based on passed boolen, it display's full info or just message with the error level.
func NewAppLogger(mode string) *AppLogger {
	l := &AppLogger{
		logger: slog.New(loggerHandler(os.Stdout)),
	}
	l.SetMode(mode)
	return l

}

// Error function extract information about error (AppErr) object.
func (l *AppLogger) Error(err error) {
	if err == nil {
		return
	}

	var appErr *apperr.AppErr
	if errors.As(err, &appErr) {
		if l.mode == "info" || l.mode == "debug" {
			l.logger.Error("error occurred",
				"message", appErr.Message,
				"code", appErr.Code,
				"origin", appErr.Origin,
				"status", appErr.StatusCode,
				"underlying", fmt.Sprintf("%+v", appErr.Err),
			)
		} else {
			l.logger.Error(appErr.Message)
		}
		return
	}

	l.logger.Error("error occurred", "error", fmt.Sprintf("%+v", err))
}

// Info function for info level messages.
// Visible only in info or debug modes.
func (l *AppLogger) Info(msg string, attrs ...any) {
	l.logger.Info(msg, attrs...)
}

// Debug function for debug level messages.
// Visible only in debug mode.
func (l *AppLogger) Debug(msg string, attrs ...any) {
		l.logger.Debug(msg, attrs...)
}

// SetVerbose function allows to set verbose mode after allocating the logger.
func (l *AppLogger) SetMode(mode string) {
	l.mode = mode

	switch mode {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "basic", "error":
		logLevel.Set(slog.LevelError)
	default:
		logLevel.Set(slog.LevelInfo)
	}
}
