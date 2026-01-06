// apperr package contains definition of errors.
package apperr

import "fmt"

// Origin indicates where the error comes from.
type Origin string

const (
	// errors caused by user input or misuse of the app
	OriginUser Origin = "user"
	// error caused by an internal library/package
	OriginInternal Origin = "internal"
	// error caused by external dependencies or systems, not under control.
	OriginExternal Origin = "external"
)

// AppErr custom error object.
// Provides additional information about errors and more user-friendly messages.
type AppErr struct {
	// user-friendly message
	Message string
	// e.g., HTTP code or custom app code
	StatusCode int
	// machine-readable code
	Code string
	// user | internal | external
	Origin Origin
	Err    error
}

// Unwrap function allows errors.Unwrap to retrieve the underlying error.
func (e *AppErr) Unwrap() error {
	return e.Err
}

// Error function satisfies error interface.
func (e *AppErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// TODO: I guess, better to make StatusCode as a enum too,
// this way i will get consistency and more eassily way to identify them.
