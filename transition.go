package gostateMachine

import (
	"errors"
	"strings"
)

type CallBack interface {
	BeforeRunCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
	RunEffectCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
	AffterRunCallBack(runableState State, data interface{}, metaData map[string]interface{}) error
}

type callBackBlank struct{}

func (*callBackBlank) BeforeRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}
func (*callBackBlank) RunEffectCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return errors.New("callBackBlank is default,plase reset it")
}
func (*callBackBlank) AffterRunCallBack(s State, data interface{}, metaData map[string]interface{}) error {
	return nil
}

//定义过渡器
type Transition struct {
	startState   State
	nextState    State
	runableState State
	name         string
	callBack     string
	stateMachine *StateMachine
	inputs       []interface{}
	metaData     map[string]interface{}
}

func NewTransition(transitionName string, startState, nextState State, sm *StateMachine) *Transition {
	transitionName = strings.ToUpper(transitionName)
	return &Transition{
		name:         transitionName,
		startState:   startState,
		nextState:    nextState,
		stateMachine: sm,
		inputs:       make([]interface{}, 0),
		metaData:     make(map[string]interface{}),
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

func (t *Transition) PutMetaData(key string, data interface{}) {
	t.metaData[key] = data
}

func (t *Transition) GetMetaDataByKey(key string) (interface{}, bool) {
	data, ok := t.metaData[key]
	return data, ok
}

func (t *Transition) SetCallBackName(name string) {
	t.callBack = name
}

func (t *Transition) Execute(data interface{}) (State, error) {
	var callBack CallBack
	if c, ok := t.stateMachine.callBacks[t.callBack]; ok {
		callBack = c
	} else {
		callBack = &callBackBlank{}
	}
	if err := callBack.BeforeRunCallBack(t.runableState, data, t.metaData); err != nil {
		return t.runableState, err
	}
	if err := callBack.RunEffectCallBack(t.runableState, data, t.metaData); err != nil {
		return t.runableState, err
	}
	t.runableState = t.nextState
	if err := callBack.AffterRunCallBack(t.runableState, data, t.metaData); err != nil {
		return t.runableState, err
	}
	return t.runableState, nil
}
