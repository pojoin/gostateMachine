package gostateMachine

type State struct{
	stateValue	string
}

func NewState(stateValue string) State{
	return State{stateValue :stateValue}
}

func (s *State) GetStateValue() string{
	return s.stateValue
}

func (s *State) SetStateValue(stateValue string){
	s.stateValue = stateValue
}

func (s *State) IsEquels(state State) bool {
	if s.stateValue == state.GetStateValue(){
		return true
	}else {
		return false
	}
}
