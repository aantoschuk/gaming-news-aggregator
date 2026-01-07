package root

import (
	"cmp"

	"github.com/spf13/cobra"
)

type flags struct {
	v bool
}

func retrieveFlags(cmd *cobra.Command) (flags, error) {
	v, errCannotGetVFlag := cmd.Flags().GetBool("verbose")
	var f flags
	if err := cmp.Or(errCannotGetVFlag); err != nil {
		return f, err

	}
	f.v = v
	return f, nil

}
