package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type Fetch Interfaces.Action

func (p Fetch) GetParams(params []string) []string {
	return Utils.Prepend("fetch", params)
}

func (p Fetch) GetActionName() string {
	return "Fetch"
}

func NewFetch() Interfaces.IAction {
	return Fetch{}
}
