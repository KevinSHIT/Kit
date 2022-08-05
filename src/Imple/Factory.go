package Imple

import "github.com/KevinZonda/Kit/src/Interfaces"

type ModuleType int

const (
	BaseModule ModuleType = iota
	PushModule
	CommitModule
	ForcePushModule
	MergeModule
)

func Factory(m ModuleType) Interfaces.IAction {
	switch m {
	case PushModule:
		return NewPush()
	case BaseModule:
		return NewBase()
	case CommitModule:
		return NewCommit()
	case ForcePushModule:
		return NewForcePush()
	case MergeModule:
		return NewMerge()
	}
	return nil
}
