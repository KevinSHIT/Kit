package Imple

import (
	"fmt"
	"github.com/KevinZonda/Kit/src/Utils"
)

func Init(params []string) {
	switch len(params) {
	case 2:
		initCmd(params[0], params[1], false, "")
	case 3:
		initCmd(params[0], params[1], true, params[2])
	default:
		getInitHelp()
	}
}

func initCmd(targetGit string, branchName string, autoAdd bool, commitMsg string) {
	Utils.RunCmd("git", "init")
	Utils.RunCmd("git", "add", ".")
	if autoAdd {
		Utils.RunCmd("git", "commit", "-m", commitMsg)
	}
	Utils.RunCmd("git", "branch", "-M", branchName)
	Utils.RunCmd("git", "remote", "add", "origin", targetGit)
	if autoAdd {
		Utils.RunCmd("git", "push", "-u", "origin", branchName)
	}
}

func getInitHelp() {
	fmt.Println("Usage: kit init <targetGit> <branchName> [commitMsg]")
	fmt.Println("       if contains commitMsg, will add and push auto")
}
