package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type Push Interfaces.Action

func (p Push) GetParams(params []string) []string {
	switch len(params) {
	case 1:
		if params[0] == "-f" {
			return []string{"push", "--force"}
		}
		return []string{"push", "origin", params[0]}
	case 2:
		if params[0] == "-f" {
			return []string{"push", "--force", params[1]}
		}
	}
	return Utils.Prepend("push", params)
}

func (p Push) GetActionName() string {
	return "Push"
}

func NewPush() Interfaces.IAction {
	return Push{}
}

func (p Push) GetHelp() string {
	return "push [-f] [branch]"
}
