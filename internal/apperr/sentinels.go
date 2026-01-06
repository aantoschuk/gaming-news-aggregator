package apperr

var (
	ErrMissingRequiredFlag = &AppErr{
		Message:    "required flags are missing",
		StatusCode: 1,
		Code:       "MISSING_REQ_FLAGS",
		Origin:     OriginUser,
	}
)
