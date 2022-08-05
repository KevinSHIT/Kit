package Imple

import "github.com/KevinZonda/Kit/src/Utils"

func Acp(params []string) {
	acp(params[0])
}

func acp(msg string) {
	Utils.RunCmd("git", "add", ".")
	Utils.RunCmd("git", "commit", "-m", msg)
	Utils.RunCmd("git", "push")
}
