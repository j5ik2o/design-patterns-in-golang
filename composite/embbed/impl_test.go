package embbed

import "testing"

func TestComposite(t *testing.T) {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	binDir = binDir.Add(NewFile("vi", 10000).Entry)
	binDir = binDir.Add(NewFile("latex", 20000).Entry)

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	yuki = yuki.Add(NewFile("diary.html", 100).Entry)
	yuki = yuki.Add(NewFile("Composite.java", 200).Entry)

	hanako = hanako.Add(NewFile("memo.tex", 300).Entry)

	tomura = tomura.Add(NewFile("game.doc", 400).Entry)
	tomura = tomura.Add(NewFile("junk.mail", 500).Entry)

	usrDir = usrDir.Add(yuki.Entry)
	usrDir = usrDir.Add(hanako.Entry)
	usrDir = usrDir.Add(tomura.Entry)

	rootDir = rootDir.Add(binDir.Entry)
	rootDir = rootDir.Add(tmpDir.Entry)
	rootDir = rootDir.Add(usrDir.Entry)

	rootDir.PrintList()
}
