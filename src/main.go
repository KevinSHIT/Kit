package main

import (
	"fmt"
	"github.com/KevinZonda/Kit/src/Imple"
	"github.com/KevinZonda/Kit/src/Interfaces"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var act Interfaces.IAction = nil
	if len(args) == 0 {
		act = Imple.Factory(Imple.BaseModule)
	} else {
		actArg := args[0]
		args = args[1:]
		switch strings.ToLower(actArg) {
		case "p":
			act = Imple.Factory(Imple.PushModule)
		case "c":
			act = Imple.Factory(Imple.CommitModule)
		case "fp":
			act = Imple.Factory(Imple.ForcePushModule)
		case "m":
			act = Imple.Factory(Imple.MergeModule)
		case "i", "new":
			PrintModuleName("Init")
			Imple.Init(args)
		case "acp":
			PrintModuleName("Acp")
			Imple.Acp(args)
		default:
			act = Imple.Factory(Imple.BaseModule)
		}
	}
	if act == nil {
		return
	}
	PrintModuleName(act.GetActionName())
	Interfaces.Do(act, args)
}

func PrintModuleName(module string) {
	fmt.Printf("[Kit by KevinZonda - %s]\n", module)
}
