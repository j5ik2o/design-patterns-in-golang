package state

import "fmt"

type State interface {
	fmt.Stringer
	DoClock(context Context, hour uint)
	DoUse(context Context)
	DoAlarm(context Context)
	DoPhone(context Context)
}
