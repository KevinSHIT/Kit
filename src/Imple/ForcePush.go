package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
)

type ForcePush Interfaces.Action

func (p ForcePush) GetParams(params []string) []string {
	return append([]string{"push", "--force"}, params...)
}

func (p ForcePush) GetActionName() string {
	return "ForcePush"
}

func NewForcePush() Interfaces.IAction {
	return ForcePush{}
}
