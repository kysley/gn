package main

import (
	"os/exec"
	"runtime"
)

func sleep() *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("powershell", "shutdown /h")
	} else if runtime.GOOS == "darwin" {
		return exec.Command("bash", "systemctl hibernate")
	}

	panic("OS not supported")
}
