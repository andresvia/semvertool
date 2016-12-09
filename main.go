// Simple tool to compare desired semver range vs actual version
// Usage: semvertool <desired_range> <actual_version>
// Returns 0 if the actual_version is inside the desired_range
// otherwise returns 1.
package main

import (
	"flag"
	"github.com/blang/semver"
	"os"
)

func main() {
	flag.Parse()
	desired_range := semver.MustParseRange(flag.Args()[0])
	actual_version := semver.MustParse(flag.Args()[1])
	if desired_range(actual_version) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
