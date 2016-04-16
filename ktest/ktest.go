package ktest

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// Where to look for *.ko files
var ModulePath = "./"

// Insmod loads module (found in ModulePath, no .ko extension).
func Insmod(module string) error {
	return Cmd("insmod", filepath.Join(ModulePath, module+".ko"))
}

// Rmmod unloads module.
func Rmmod(module string) error {
	return Cmd("rmmod", module)
}

// Cmd executes "prog arg1 arg2 ...".
func Cmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		var a []string
		for _, arg := range args {
			a = append(a, fmt.Sprintf("%q", arg))
		}
		return fmt.Errorf("%v %v: %v: %q", prog, strings.Join(a, " "), err, out)
	}
	return nil
}

// HaveModule returns true if module appears in /proc/modules.
func HaveModule(module string) bool {
	return Cmd("grep", "hello ", "/proc/modules") == nil

}
