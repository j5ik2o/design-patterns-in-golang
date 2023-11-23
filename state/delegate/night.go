package delegate

type Night struct {
}

func (n *Night) String() string {
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
