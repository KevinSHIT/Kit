package main

import (
	"fmt"
	"github.com/KevinZonda/Kit/src/Imple"
	"github.com/KevinZonda/Kit/src/Interfaces"
	"os"
)

func main() {
	args := os.Args[1:]
	var act Interfaces.IAction
	if len(args) == 0 {
		act = Imple.Factory(Imple.BaseModule)
	} else {
		actArg := args[0]
		args = args[1:]
		switch actArg {
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
			return
		default:
			act = Imple.Factory(Imple.BaseModule)
		}
	}
	PrintModuleName(act.GetActionName())
	Interfaces.Do(act, args)
}

func PrintModuleName(module string) {
	fmt.Printf("[Kit by KevinZonda - %s]\n", module)
}
