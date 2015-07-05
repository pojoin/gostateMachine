package gostateMachine

//状态机定义
type StateMachine struct{
	Name string
	transitionMap	map[string]*Transition
}


func NewStateMachine(name string) StateMachine{
	return StateMachine{Name : name}
}


