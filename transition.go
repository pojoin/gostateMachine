package gostateMachine

import "strings"

type CallBack interface {
	BeforeRunCallBack(runableState State, data interface{})
	RunEffectCallBack(runableState State, data interface{}) bool
	AffterRunCallBack(runableState State, data interface{})
}

type callBackBlank struct{}

func (*callBackBlank) BeforeRunCallBack(s State, data interface{}) {}
func (*callBackBlank) RunEffectCallBack(s State, data interface{}) bool {
	return true
}
func (*callBackBlank) AffterRunCallBack(s State, data interface{}) {}

//定义过渡器
type Transition struct {
	startState   State
	nextState    State
	runableState State
	name         string
	callBack     string
	stateMachine *StateMachine
	inputs       []interface{}
}

func NewTransition(transitionName string, startState, nextState State, sm *StateMachine) *Transition {
	transitionName = strings.ToUpper(transitionName)
	return &Transition{
		name:         transitionName,
		startState:   startState,
		nextState:    nextState,
		stateMachine: sm,
		inputs:       make([]interface{}, 0),
	}
}

func (t *Transition) Name() string {
	return t.name
}

func (t *Transition) AppendInput(input interface{}) {
	t.inputs = append(t.inputs, input)
}

func (t *Transition) GetInputs() []interface{} {
	return t.inputs
}

func (t *Transition) Execute(data interface{}) State {
	var callBack CallBack
	if c, ok := t.stateMachine.callBacks[t.callBack]; ok {
		callBack = c
	} else {
		callBack = &callBackBlank{}
	}
	callBack.BeforeRunCallBack(t.runableState, data)
	if !callBack.RunEffectCallBack(t.runableState, data) {
		return t.runableState
	}
	t.runableState = t.nextState
	callBack.AffterRunCallBack(t.runableState, data)
	return t.runableState
}
