package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type Pull Interfaces.Action

func (p Pull) GetParams(params []string) []string {
	return Utils.Prepend("pull", params)
}

func (p Pull) GetActionName() string {
	return "Pull"
}

func NewPull() Interfaces.IAction {
	return Pull{}
}
