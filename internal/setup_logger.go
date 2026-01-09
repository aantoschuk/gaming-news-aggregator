package internal

import "github.com/aantoschuk/feed/internal/app_logger"

// SetupLogger function accepts mode flags and returns a logger
// based on this flags.
func SetupLogger(d, v bool) *app_logger.AppLogger {
	mode := "basic"
	logger := app_logger.NewAppLogger(mode)

	if d {
		logger.SetMode("debug")
	} else if v {
		logger.SetMode("info")
	}

	return logger
}
