package internal

import "flag"

// HandleFlags function is create a flags and derefences them.
func HandleFlags() (bool, bool) {
	// local flags
	v := flag.Bool("v", false, "enable verbose mode output")
	d := flag.Bool("d", false, "enable debug mode output")
	flag.Parse()
	return *d, *v
}
