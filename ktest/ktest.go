package ktest

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Where to look for *.ko files
var ModulePath = "./"

// MustInsmod loads module (found in ModulePath, no .ko extension),
// and fails t in case of an error
func MustInsmod(t *testing.T, module string) {
	MustCmd(t, "insmod", filepath.Join(ModulePath, module+".ko"))
}

// MustCmd executes the command "prog arg1 arg2 ...",
// and fails t in case of an error.
func MustCmd(t *testing.T, prog string, args ...string) {
	cmd := exec.Command(prog, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		var a []string
		for _, arg := range args {
			a = append(a, fmt.Sprintf("%q", arg))
		}
		t.Fatalf("%v %v: %v\n%q", prog, strings.Join(a, " "), err, out)
	}
}
