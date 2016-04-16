package ktest

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
	"testing"
)

var ModulePath = "./"

func MustInsmod(t *testing.T, fname string) {
	MustCmd(t, "insmod", path.Join(ModulePath, fname, ".ko"))
}

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
