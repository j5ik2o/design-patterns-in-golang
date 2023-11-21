package embed

type CountDisplay struct {
	*DisplayBase
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{NewDisplay(impl)}
}

func (c *CountDisplay) MultiDisplay(times int) {
	c.Open()
	for i := 0; i < times; i++ {
		c.Print()
	}
	c.Close()
}
