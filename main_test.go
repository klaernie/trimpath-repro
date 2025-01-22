package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestRun(t *testing.T) {
	/*
		why care about the path separators?

		The absolute paths in the testdata comes out differently, depending on the path separator
		used by the system that runs the tests.

		This leads to the problem, that one cannot verify a snapshot taken under Linux when using Windows
		and vice versa, unless the tests are differentiated by the path separator.
	*/
	var safePathSep string
	if os.PathSeparator == '/' {
		safePathSep = "slash"
	} else if os.PathSeparator == '\\' {
		safePathSep = "backslash"
	} else {
		t.Fatal("current path separator is unexpected - please fix test to handle this path separator")
	}
	s := snaps.WithConfig(
		snaps.Dir(filepath.Join("__snapshots__", "pathseparator-"+safePathSep)),
	)

	s.MatchSnapshot(t, run())
}
