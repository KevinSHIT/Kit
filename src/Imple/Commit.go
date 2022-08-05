package Imple

import (
	"github.com/KevinZonda/Kit/src/Interfaces"
	"github.com/KevinZonda/Kit/src/Utils"
)

type Commit Interfaces.Action

func (c Commit) GetParams(params []string) []string {
	switch len(params) {
	case 0:
		return []string{"commit", "-a ", "--allow-empty-message", "-m", "''"}
	case 1:
		return []string{"commit", "-m", params[0]}
	}
	return Utils.Prepend("commit", params)
}

func (c Commit) GetActionName() string {
	return "Commit"
}

func NewCommit() Interfaces.IAction {
	return Commit{}
}
