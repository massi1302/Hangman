package Hangman

import (
	"os"
	"os/exec"
	"runtime"
)

var clearFuncs map[string]func()

func init() {
	clearFuncs = make(map[string]func())
	clearFuncs["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearConsole() {
	value, ok := clearFuncs[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clearFuncs terminal screen :(")
	}
}
