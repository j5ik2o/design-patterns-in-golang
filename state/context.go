package state

type Context interface {
	SetClock(hour uint)
	ChangeState(state State)
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}
