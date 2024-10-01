package Hangman

import (
	"os"
	"os/exec"
	"runtime"
)

var clearFuncs map[string]func()

func init() {
	clearFuncs = make(map[string]func()) //Initialize it
	clearFuncs["linux"] = func() {
		cmd := exec.Command("clearFuncs") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
	clearFuncs["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
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
