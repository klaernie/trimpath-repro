package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestRun(t *testing.T) {
	snaps.MatchSnapshot(t, run())
}
