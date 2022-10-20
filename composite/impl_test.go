package composite

import "testing"

func TestComposite(t *testing.T) {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	rootDir.Add(binDir)
	rootDir.Add(tmpDir)
	rootDir.Add(usrDir)

	binDir.Add(NewFile("vi", 10000))
	binDir.Add(NewFile("latex", 20000))

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	usrDir.Add(yuki)
	usrDir.Add(hanako)
	usrDir.Add(tomura)

	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.java", 200))

	hanako.Add(NewFile("memo.tex", 300))

	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))

	rootDir.PrintLine()
}
