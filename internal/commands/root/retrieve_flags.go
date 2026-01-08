package root

import (
	"cmp"

	"github.com/spf13/cobra"
)

type flags struct {
	v bool
	d bool
}

func retrieveFlags(cmd *cobra.Command) (flags, error) {
	v, errCannotGetVFlag := cmd.Flags().GetBool("verbose")
	d, errCannotGetVFlag := cmd.Flags().GetBool("debug")
	var f flags
	if err := cmp.Or(errCannotGetVFlag); err != nil {
		return f, err

	}
	f.v = v
	f.d = d
	return f, nil

}
