package chain_of_responsibility

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) *Trouble {
	return &Trouble{
		number: number,
	}
}

func (t *Trouble) GetNumber() int {
	return t.number
}

func (t *Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.number)
}

type SupportBase interface {
	Done(trouble *Trouble)
	Fail(trouble *Trouble)
	Next(trouble *Trouble) Support
}

type Support interface {
	Resolve(trouble *Trouble) bool
	Support(trouble *Trouble)
}

// ---

type NoSupport struct {
	name string
	next Support
}

func NewNoSupport(name string, next Support) *NoSupport {
	return &NoSupport{
		name,
		next,
	}
}

func (n *NoSupport) ToString() string {
	return fmt.Sprintf("[%s@NoSupport]", n.name)
}

func (n *NoSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, n.ToString())
}

func (n *NoSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (n *NoSupport) Next() Support {
	return n.next
}

func (n *NoSupport) Resolve(trouble *Trouble) bool {
	return false
}

func (n *NoSupport) Support(trouble *Trouble) {
	if n.Resolve(trouble) {
		n.Done(trouble)
	} else if n.Next() != nil {
		n.Next().Support(trouble)
	} else {
		n.Fail(trouble)
	}
}

// ---

type LimitSupport struct {
	name  string
	next  Support
	limit int
}

func NewLimitSupport(name string, limit int, next Support) *LimitSupport {
	return &LimitSupport{
		name,
		next,
		limit,
	}
}

func (l *LimitSupport) ToString() string {
	return fmt.Sprintf("[%s@LimitSupport]", l.name)
}

func (l *LimitSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, l.ToString())
}

func (l *LimitSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (l *LimitSupport) Next() Support {
	return l.next
}

func (l *LimitSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber() < l.limit {
		return true
	}
	return false
}

func (l *LimitSupport) Support(trouble *Trouble) {
	if l.Resolve(trouble) {
		l.Done(trouble)
	} else if l.Next() != nil {
		l.Next().Support(trouble)
	} else {
		l.Fail(trouble)
	}
}

// ---

type OddSupport struct {
	name string
	next Support
}

func NewOddSupport(name string, next Support) *OddSupport {
	return &OddSupport{
		name,
		next,
	}
}

func (o *OddSupport) ToString() string {
	return fmt.Sprintf("[%s@OddSupport]", o.name)
}

func (o *OddSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, o.ToString())
}

func (o *OddSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (o *OddSupport) Next() Support {
	return o.next
}

func (o *OddSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber()%2 == 1 {
		return true
	}
	return false
}

func (o *OddSupport) Support(trouble *Trouble) {
	if o.Resolve(trouble) {
		o.Done(trouble)
	} else if o.Next() != nil {
		o.Next().Support(trouble)
	} else {
		o.Fail(trouble)
	}
}

// ---

type SpecialSupport struct {
	name   string
	next   Support
	number int
}

func NewSpecialSupport(name string, number int, next Support) *SpecialSupport {
	return &SpecialSupport{
		name,
		next,
		number,
	}
}

func (s *SpecialSupport) ToString() string {
	return fmt.Sprintf("[%s@SpecialSupport]", s.name)
}

func (s *SpecialSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, s.ToString())
}

func (s *SpecialSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (s *SpecialSupport) Next() Support {
	return s.next
}

func (s *SpecialSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber() == s.number {
		return true
	}
	return false
}

func (s *SpecialSupport) Support(trouble *Trouble) {
	if s.Resolve(trouble) {
		s.Done(trouble)
	} else if s.Next() != nil {
		s.Next().Support(trouble)
	} else {
		s.Fail(trouble)
	}
}
