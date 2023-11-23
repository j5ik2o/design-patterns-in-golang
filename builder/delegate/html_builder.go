package delegate

import (
	"os"
	"path/filepath"
	"strings"
)

type HTMLBuilder struct {
	workDir       string
	filename      string
	stringBuilder strings.Builder
}

func NewHTMLBuilder(workDir string) *HTMLBuilder {
	return &HTMLBuilder{
		workDir:  workDir,
		filename: "untitled.html",
	}
}

func (b *HTMLBuilder) MakeTitle(title string) {
	b.filename = title + ".html"
	b.stringBuilder.WriteString("<!DOCTYPE html>\n")
	b.stringBuilder.WriteString("<html>\n")
	b.stringBuilder.WriteString("<head><title>")
	b.stringBuilder.WriteString(title)
	b.stringBuilder.WriteString("</title></head>\n")
	b.stringBuilder.WriteString("<body>\n")
	b.stringBuilder.WriteString("<h1>")
	b.stringBuilder.WriteString(title)
	b.stringBuilder.WriteString("</h1>\n\n")
}

func (b *HTMLBuilder) MakeString(str string) {
	b.stringBuilder.WriteString("<p>")
	b.stringBuilder.WriteString(str)
	b.stringBuilder.WriteString("</p>\n\n")
}

func (b *HTMLBuilder) MakeItems(items []string) {
	b.stringBuilder.WriteString("<ul>\n")
	for _, s := range items {
		b.stringBuilder.WriteString("<li>")
		b.stringBuilder.WriteString(s)
		b.stringBuilder.WriteString("</li>\n")
	}
	b.stringBuilder.WriteString("</ul>\n\n")
}

func (b *HTMLBuilder) Close() error {
	b.stringBuilder.WriteString("</body>")
	b.stringBuilder.WriteString("</html>\n")
	path := filepath.Join(b.workDir, b.filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(b.stringBuilder.String())
	return err
}

func (b *HTMLBuilder) GetHTMLResult() string {
	return filepath.Join(b.workDir, b.filename)
}
