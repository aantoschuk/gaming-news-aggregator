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

// AppLogger object. Based on the slog with additional parameter
// to handle amount of visible  information in the console.
type AppLogger struct {
	// slog
	logger *slog.Logger
	// controls amount of visible information
	verbose bool
}

func loggerHandler(w io.Writer) slog.Handler {
	return slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug,
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
func NewAppLogger(v bool) *AppLogger {
	var handler slog.Handler
	if v {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = loggerHandler(os.Stdout)
	}

	return &AppLogger{
		logger:  slog.New(handler),
		verbose: v,
	}
}

// Error function extract information about error (AppErr) object.
func (l *AppLogger) Error(err error) {
	if err == nil {
		return
	}

	// extract apprr info
	var appErr *apperr.AppErr
	if errors.As(err, &appErr) {
		if l.verbose {
			// Full verbose structured log
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

	// Not an AppErr
	if l.verbose {
		l.logger.Error("error occurred", "error", fmt.Sprintf("%+v", err))
	} else {
		l.logger.Error(err.Error())
	}
}

// Info function for info level messages.
// Visible only in verbose mode.
func (l *AppLogger) Info(msg string, attrs ...any) {
	if l.verbose {
		l.logger.Info(msg, attrs...)
	}
}

// Debug function for debug level messages.
// Visible only in verbose mode.
func (l *AppLogger) Debug(msg string, attrs ...any) {
	if l.verbose {
		l.logger.Debug(msg, attrs...)
	}
}

// SetVerbose function allows to set verbose mode after allocating the logger.
func (l *AppLogger) SetVerbose(v bool) {
	l.verbose = v
}
