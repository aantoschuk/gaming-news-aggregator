package apperr

import (
	"errors"
	"fmt"
	"testing"
)

// testing basic usability of AppErr
func TestAppErr_Basic(t *testing.T) {
	underlying := fmt.Errorf("underlying error")
	appErr := &AppErr{
		Message:    "user-friendly message",
		Code:       "MY_CODE",
		Origin:     OriginInternal,
		StatusCode: 500,
		Err:        underlying,
	}

	t.Run("Error returns correct string", func(t *testing.T) {
		expected := "user-friendly message: underlying error"
		if appErr.Error() != expected {
			t.Fatalf("expected %q, got %q", expected, appErr.Error())
		}
	})

	t.Run("Unwrap returns underlying error", func(t *testing.T) {
		if errors.Unwrap(appErr) != underlying {
			t.Fatal("Unwrap did not return the underlying error")
		}
	})

	t.Run("Fields are correctly set", func(t *testing.T) {
		if appErr.Code != "MY_CODE" {
			t.Fatal("Code field not set correctly")
		}
		if appErr.Origin != OriginInternal {
			t.Fatal("Origin field not set correctly")
		}
		if appErr.StatusCode != 500 {
			t.Fatal("StatusCode field not set correctly")
		}
	})
}

var ErrInvalidURL = &AppErr{Code: "INVALID_URL"}

// TestAppErr_IsAndAs function tests AppErr for errors.Is and errors.As compability
func TestAppErr_IsAndAs(t *testing.T) {
	wrapped := fmt.Errorf("wrapped: %w", ErrInvalidURL)

	if !errors.Is(wrapped, ErrInvalidURL) {
		t.Fatal("errors.Is failed to detect the sentinel error")
	}

	var e *AppErr
	if !errors.As(wrapped, &e) {
		t.Fatal("errors.As failed to extract *AppErr")
	}
	if e.Code != "INVALID_URL" {
		t.Fatal("errors.As returned wrong AppErr")
	}
}

// TestAppErr_Constructors tests function defined in constructors.go
func TestAppErr_Constructors(t *testing.T) {
	err := NewInternalError("fail", "INTERNAL_FAIL", 500, fmt.Errorf("db error"))

	if err.Origin != OriginInternal {
		t.Fatal("Origin not set correctly")
	}
	if err.Code != "INTERNAL_FAIL" {
		t.Fatal("Code not set correctly")
	}
	if err.StatusCode != 500 {
		t.Fatal("StatusCode not set correctly")
	}
	if errors.Unwrap(err) == nil {
		t.Fatal("underlying error not set")
	}
}
