package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type CherryPick Interfaces.Action

func (c CherryPick) GetParams(params []string) []string {
	return Utils.Prepend("cherry-pick", params)
}

func (c CherryPick) GetActionName() string {
	return "CherryPick"
}

func NewCherryPick() Interfaces.IAction {
	return CherryPick{}
}
