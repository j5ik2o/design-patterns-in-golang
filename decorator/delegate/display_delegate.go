package delegate

type DisplayDelegate interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) string
}
