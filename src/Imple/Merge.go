package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type Merge Interfaces.Action

func (b Merge) GetParams(params []string) []string {
	return Utils.Prepend("merge", params)
}

func (b Merge) GetActionName() string {
	return "Merge"
}

func NewMerge() Interfaces.IAction {
	return Merge{}
}
