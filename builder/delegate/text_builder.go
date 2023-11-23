package delegate

import "strings"

type TextBuilder struct {
	buffer strings.Builder
}

func NewTextBuilder() *TextBuilder {
	var sb strings.Builder
	sb.Grow(1024)
	return &TextBuilder{sb}
}

func (t *TextBuilder) MakeTitle(title string) {
	t.buffer.WriteString("==============================\n")
	t.buffer.WriteString("[" + title + "]\n")
	t.buffer.WriteString("\n")
}

func (t *TextBuilder) MakeString(str string) {
	t.buffer.WriteString("■" + str + "\n")
	t.buffer.WriteString("\n")
}

func (t *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		t.buffer.WriteString("　・" + item + "\n")
	}
	t.buffer.WriteString("\n")
}

func (t *TextBuilder) Close() error {
	t.buffer.WriteString("==============================\n")
	return nil
}

func (t *TextBuilder) GetTextResult() string {
	return t.buffer.String()
}
