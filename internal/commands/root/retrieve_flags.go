package root

import (
	"cmp"

	"github.com/aantoschuk/feed/internal/apperr"
	"github.com/spf13/cobra"
)

type flags struct {
	v bool
	u string
}

func retrieveFlags(cmd *cobra.Command) (flags, error) {
	v, errCannotGetVFlag := cmd.Flags().GetBool("verbose")
	u, errCannotGetUFlag := cmd.Flags().GetString("url")
	var f flags
	if errCannotGetUFlag != nil {
		appErr := apperr.NewInternalError("cannot retrieve -u flag", "RETRIEVE_U_FLAG_EROR", 1, errCannotGetUFlag)
		return f, appErr
	}
	if err := cmp.Or(errCannotGetVFlag, errCannotGetUFlag); err != nil {
		return f, err

	}
	if u == "" {
		return f, apperr.ErrMissingRequiredFlag
	}
	f.u = u
	f.v = v
	return f, nil

}
