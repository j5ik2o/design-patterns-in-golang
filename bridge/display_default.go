package bridge

import "desgin-patterns-in-golang/bridge/delegates"

type DisplayDefault struct {
	underlying delegates.DisplayDelegate
}

func NewDisplayDefault(underlying delegates.DisplayDelegate) *DisplayDefault {
	return &DisplayDefault{underlying}
}

func (d *DisplayDefault) Open() {
	d.underlying.RawOpen()
}

func (d *DisplayDefault) Print() {
	d.underlying.RawPrint()
}

func (d *DisplayDefault) Close() {
	d.underlying.RawClose()
}

func (d *DisplayDefault) Display() {
	d.Open()
	d.Print()
	d.Close()
}
