package app_logger

import (
	"bytes"
	"fmt"
	"log/slog"
	"strings"
	"testing"

	"github.com/aantoschuk/feed/internal/apperr"
)

func newTestLogger(mode string) (*AppLogger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	l := &AppLogger{
		logger: slog.New(loggerHandler(buf)),
	}

	l.SetMode(mode)
	return l, buf
}

func TestAppLogger_Error_AppErr(t *testing.T) {
	appErr := &apperr.AppErr{
		Message:    "something went wrong",
		Code:       "TEST_CODE",
		Origin:     apperr.OriginInternal,
		StatusCode: 500,
		Err:        fmt.Errorf("underlying error"),
	}

	logger, buf := newTestLogger("info")
	logger.Error(appErr)

	out := buf.String()
	if !strings.Contains(out, "something went wrong") ||
		!strings.Contains(out, "TEST_CODE") ||
		!strings.Contains(out, "underlying error") {
		t.Fatal("verbose AppErr logging failed")
	}
}

func TestAppLogger_Error_AppErr_NonVerbose(t *testing.T) {
	appErr := &apperr.AppErr{
		Message:    "something went wrong",
		Code:       "TEST_CODE",
		Origin:     apperr.OriginInternal,
		StatusCode: 500,
		Err:        fmt.Errorf("underlying error"),
	}

	logger, buf := newTestLogger("basic")
	logger.Error(appErr)

	out := buf.String()
	if !strings.Contains(out, "something went wrong") || strings.Contains(out, "TEST_CODE") {
		t.Fatal("non-verbose AppErr logging failed")
	}
}

func TestAppLogger_Error_NonAppErr(t *testing.T) {
	err := fmt.Errorf("simple error")
	logger, buf := newTestLogger("basic")
	logger.Error(err)

	out := buf.String()
	if !strings.Contains(out, "simple error") {
		t.Fatal("non-AppErr verbose logging failed")
	}
}

func TestAppLogger_SetVerbose(t *testing.T) {
	logger, buf := newTestLogger("basic")
	if logger.mode != "basic" {
		t.Fatal("expected verbose=false initially")
	}

	logger.SetMode("info")
	if logger.mode != "info" {
		t.Fatal("SetVerbose did not update verbose flag")
	}
	appErr := &apperr.AppErr{Message: "error!"}
	logger.Error(appErr)

	out := buf.String()
	if !strings.Contains(out, "error!") {
		t.Fatal("SetVerbose did not enable verbose logging")
	}
}

func TestAppLogger_InfoDebug(t *testing.T) {
	logger, buf := newTestLogger("error")
	logger.Info("info message")
	logger.Debug("debug message")
	if buf.Len() != 0 {
		t.Fatal("non-verbose Info/Debug should not log")
	}

	logger.SetMode("debug")
	logger.Info("info message")
	logger.Debug("debug message")
	out := buf.String()

	t.Log("out")
	t.Log(out)
	if !strings.Contains(out, "info message") || !strings.Contains(out, "debug message") {
		t.Fatalf("verbose Info/Debug logging failed, got: %s", out)
	}
}
