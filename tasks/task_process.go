package tasks

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"eaciit/diagnostic"
)

//IsProcessRunning Pre-defined task that is used for checking wether a process
//is currently running or not
func IsProcessRunning(process string) diagnostic.Task {
	var out, err bytes.Buffer
	var e error
	var cmd *exec.Cmd

	return diagnostic.TaskFunc(func() error {
		os := runtime.GOOS
		switch os {
		case "windows":
			cmd = exec.Command("cmd", "/C", "tasklist", "/fo", "csv", "/nh")
		case "linux":
			cmd = exec.Command("ps", "-def")
		default:
			cmd = nil
		}

		if cmd != nil {
			cmd.Stdout = &out
			cmd.Stderr = &err
			e = cmd.Run()

			if e != nil {
				return errors.New(err.String())
			}

			if strings.Contains(out.String(), process) {
				return nil
			}

			return errors.New(fmt.Sprint("No process named '", process, "' is running"))
		}

		return errors.New(fmt.Sprint("No OS '", os, "' command handler defined"))
	})
}
