package Interfaces

type Action struct {
	IAction
}

func (a Action) GetActionName() string {
	return "Action"
}
