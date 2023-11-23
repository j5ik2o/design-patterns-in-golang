package delegate

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(b *Banner) *PrintBanner {
	return &PrintBanner{b}
}

func (pb *PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
