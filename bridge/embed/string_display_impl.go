package embed

type StringDisplayImpl struct {
	string string
	width  int
}

func NewStringDisplayImpl(string string) *StringDisplayImpl {
	return &StringDisplayImpl{string, len(string)}
}

func (s *StringDisplayImpl) RawOpen() {
	s.printLine()
}

func (s *StringDisplayImpl) RawPrint() {
	println("|" + s.string + "|")
}

func (s *StringDisplayImpl) RawClose() {
	s.printLine()
}

func (s *StringDisplayImpl) printLine() {
	print("+")
	for i := 0; i < s.width; i++ {
		print("-")
	}
	println("+")
}
