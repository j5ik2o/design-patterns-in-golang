package flyweight

type BigCharFactory struct {
	pool map[string]*BigChar
}

func NewBigCharFactory() *BigCharFactory {
	return &BigCharFactory{
		pool: make(map[string]*BigChar),
	}
}

func (f *BigCharFactory) GetBigChar(charName string) *BigChar {
	bc, ok := f.pool[charName]
	if !ok {
		bc = NewBigChar(charName)
		f.pool[charName] = bc
	}
	return bc
}
