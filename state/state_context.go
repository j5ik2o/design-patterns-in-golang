package state

import "fmt"

type StateContext struct {
	state State
}

func NewStateContext(state State) *StateContext {
	return &StateContext{state}
}

func (sc *StateContext) Run() {
	var i = uint(0)
	for hour := uint(0); hour <= 24; hour++ {
		sc.SetClock(hour)
		switch i {
		case 0:
			sc.state.DoUse(sc)
			i = 1
		case 1:
			sc.state.DoAlarm(sc)
			i = 2
		case 2:
			sc.state.DoPhone(sc)
			i = 0
		default:
			panic("iae")
		}
	}

}

func (sc *StateContext) SetClock(hour uint) {
	sc.state.DoClock(sc, hour)
}

func (sc *StateContext) ChangeState(state State) {
	sc.state = state
}

func (sc *StateContext) CallSecurityCenter(msg string) {
	fmt.Printf("%s:%s\n", sc.state.String(), msg)
}

func (sc *StateContext) RecordLog(msg string) {
	fmt.Printf("%s:%s\n", sc.state.String(), msg)
}
