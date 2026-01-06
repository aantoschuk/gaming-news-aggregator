package apperr

// NewUserError creates error with Origin as a "user"
func NewUserError(msg string, code string, status int) *AppErr {
	return &AppErr{Message: msg, Code: code, Origin: OriginUser, StatusCode: status, Err: nil}
}

// NewUserError creates error with Origin as a "internal"
func NewInternalError(msg string, code string, status int, err error) *AppErr {
	return &AppErr{Message: msg, Code: code, Origin: OriginInternal, StatusCode: status, Err: err}
}

// NewUserError creates error with Origin as a "external"
func NewExternalError(msg string, code string, status int, err error) *AppErr {
	return &AppErr{Message: msg, Code: code, Origin: OriginExternal, StatusCode: status, Err: err}
}
