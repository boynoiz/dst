package dst_test

import (
	"testing"

	"github.com/boynoiz/dst"
)

func TestFileHeader_GoVersion(t *testing.T) {
	// This test simply asserts that the GoVersion field exists on the struct
	// and can be read/written. The full integration (copying from ast) requires
	// generation to run, but this confirms the struct update.
	f := &dst.File{
		Name:      dst.NewIdent("main"),
		GoVersion: "go1.21",
	}

	if f.GoVersion != "go1.21" {
		t.Errorf("GoVersion field not working as expected, got %q", f.GoVersion)
	}
}
