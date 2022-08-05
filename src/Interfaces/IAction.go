package Interfaces

import (
	"github.com/KevinZonda/Kit/src/Utils"
)

type IAction interface {
	GetParams(params []string) []string
	GetActionName() string
}

func Do(act IAction, params []string) {
	Utils.RunCmd("git", act.GetParams(params)...)
}
