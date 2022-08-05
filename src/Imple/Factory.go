package Imple

import "github.com/KevinZonda/Kit/src/Interfaces"

type ModuleType int

const (
	BaseModule ModuleType = iota
	PushModule
	CommitModule
	ForcePushModule
	MergeModule
	CherryPickModule
	PullModule
	FetchModule
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
	case CherryPickModule:
		return NewCherryPick()
	case PullModule:
		return NewPull()
	case FetchModule:
		return NewFetch()
	}
	return nil
}
