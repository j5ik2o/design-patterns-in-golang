package state

type Day struct {
}

func NewDay() *Day {
	return &Day{}
}

func (d *Day) String() string {
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
