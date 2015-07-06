package gostateMachine

import (
	"errors"
	"strings"
)

//状态机定义
type StateMachine struct {
	Name          string
	transitionMap map[string]*Transition
}

func NewStateMachine(name string) StateMachine {
	return StateMachine{Name: name, transitionMap: make(map[string]*Transition)}
}

func (m *StateMachine) AddTransition(tName string, startState, nextState State) {
	tName = strings.ToUpper(tName)
	t := &Transition{
		startState: startState,
		nextState:  nextState,
		Name:       tName,
	}
	m.transitionMap[tName] = t
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
		if v.startState.IsEquels(state) {
			ts = append(ts, v)
		}
	}
	return ts
}
