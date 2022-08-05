package Imple

import "github.com/KevinZonda/Kit/src/Interfaces"

type Base Interfaces.Action

func (b Base) GetParams(params []string) []string {
	return params
}

func (b Base) GetActionName() string {
	return "Base"
}

func NewBase() Interfaces.IAction {
	return Base{}
}
