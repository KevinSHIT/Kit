package main

import (
	"fmt"
	"github.com/KevinZonda/Kit/src/Imple"
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	var act Interfaces.IAction = nil
	// Magic Code, will cause KIS says I am shit
	// for _, v := range args {
	//	fmt.Printf("%s ", v)
	//}
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
		case "cp":
			act = Imple.Factory(Imple.CherryPickModule)
		case "pl":
			act = Imple.Factory(Imple.PullModule)
		case "i", "new":
			PrintModuleName("Init")
			Imple.Init(args)
		case "acp":
			PrintModuleName("Acp")
			Imple.Acp(args)

		default:
			args = Utils.Prepend(actArg, args)
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
