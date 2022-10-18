package composite

import "testing"

func TestComposite(t *testing.T) {
	rootdir := NewDirectory("root")
	bindir := NewDirectory("bin")
	tmpdir := NewDirectory("tmp")
	usrdir := NewDirectory("usr")

	rootdir.Add(bindir)
	rootdir.Add(tmpdir)
	rootdir.Add(usrdir)

	bindir.Add(NewFile("vi", 10000))
	bindir.Add(NewFile("latex", 20000))

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	usrdir.Add(yuki)
	usrdir.Add(hanako)
	usrdir.Add(tomura)

	yuki.Add(NewFile("diary.html", 100))
	yuki.Add(NewFile("Composite.java", 200))

	hanako.Add(NewFile("memo.tex", 300))

	tomura.Add(NewFile("game.doc", 400))
	tomura.Add(NewFile("junk.mail", 500))

	rootdir.PrintLine()
}
