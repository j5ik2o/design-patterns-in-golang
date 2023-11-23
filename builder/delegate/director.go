package delegate

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder,
	}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("一般的なあいさつ")
	d.builder.MakeItems([]string{
		"How are you?", "Hello.", "Hi.",
	})
	d.builder.MakeString("時間帯に応じたあいさつ")
	d.builder.MakeItems([]string{
		"Good morning.", "Good afternoon.", "Good evening.",
	})
	d.builder.Close()
}
