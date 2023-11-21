package displays

import (
	"desgin-patterns-in-golang/bridge/delegate/delegates"
)

type Display struct {
	underlying delegates.DisplayDelegate
}

func NewDisplay(underlying delegates.DisplayDelegate) *Display {
	return &Display{underlying}
}

func (d *Display) Open() {
	d.underlying.RawOpen()
}

func (d *Display) Print() {
	d.underlying.RawPrint()
}

func (d *Display) Close() {
	d.underlying.RawClose()
}

func (d *Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}
