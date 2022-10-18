package state

import "fmt"

type Context interface {
	SetClock(hour uint)
	ChangeState(state State[Context])
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}

type State[C Context] interface {
	DoClock(context C, hour uint)
	DoUse(context C)
	DoAlarm(context C)
	DoPhone(context C)
	ToString() string
}

type Day struct {
}

func (d *Day) ToString() string {
	return "[昼間]"
}

func (d *Day) DoClock(context Context, hour uint) {
	if hour < 9 || 17 <= hour {
		context.ChangeState(&Night{})
	}
}

func (d *Day) DoUse(context Context) {
	context.RecordLog("金庫使用(昼間)")
}

func (d *Day) DoAlarm(context Context) {
	context.CallSecurityCenter("非常ベル(昼間)")
}

func (d *Day) DoPhone(context Context) {
	context.RecordLog("通常の通話(昼間)")
}

type Night struct {
}

func (n *Night) ToString() string {
	return "[夜間]"
}

func (n *Night) DoClock(context Context, hour uint) {
	if 9 <= hour && hour < 17 {
		context.ChangeState(&Day{})
	}
}

func (n *Night) DoUse(context Context) {
	context.CallSecurityCenter("非常：夜間の金庫使用！")
}

func (n *Night) DoAlarm(context Context) {
	context.CallSecurityCenter("非常ベル(夜間)")
}

func (n *Night) DoPhone(context Context) {
	context.RecordLog("夜間の通話録音")
}

type StateContext struct {
	state State[Context]
}

func NewStateContext(state State[Context]) *StateContext {
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

func (sc *StateContext) ChangeState(state State[Context]) {
	sc.state = state
}

func (sc *StateContext) CallSecurityCenter(msg string) {
	fmt.Printf("%s:%s", sc.state.ToString(), msg)
}

func (sc *StateContext) RecordLog(msg string) {
	fmt.Printf("%s:%s", sc.state.ToString(), msg)
}
