package gostateMachine

import (
	"errors"
	"strings"
)

//状态机定义
type StateMachine struct {
	Name          string
	Version       string
	transitionMap map[string]*Transition
	callBacks     map[string]CallBack
	alias         map[string]string
}

func NewStateMachine(name string) StateMachine {
	return StateMachine{
		Name:          name,
		transitionMap: make(map[string]*Transition),
		callBacks:     make(map[string]CallBack),
		alias:         make(map[string]string),
	}
}

func (m *StateMachine) PutCallBacks(name string, cb CallBack) {
	m.callBacks[name] = cb
}

func (m *StateMachine) AddTransition(tName string, startState, nextState State) *StateMachine {
	tName = strings.ToUpper(tName)
	t := NewTransition(tName, startState, nextState, m)
	m.transitionMap[tName] = t
	return m
}

func (m *StateMachine) PutTransition(t *Transition) *StateMachine {
	if t.stateMachine != m {
		t.stateMachine = m
	}
	m.transitionMap[t.Name] = t
	return m
}

//获取transition
func (m *StateMachine) GetTransitionByName(tName string) (*Transition, error) {
	tName = strings.ToUpper(tName)
	if t, ok := m.transitionMap[tName]; ok {
		return t, nil
	} else {
		return nil, errors.New("can not fond transition by name" + tName)
	}
}

//获取trasintions
func (m *StateMachine) GetTransitionByState(state State) []*Transition {
	ts := make([]*Transition, 0)
	for _, v := range m.transitionMap {
		if v.startState.Equels(state) {
			ts = append(ts, v)
		}
	}
	return ts
}

func (m *StateMachine) GetAlias(key string) (value string, ok bool) {
	value, ok = m.alias[key]
	return
}
