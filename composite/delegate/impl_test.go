package delegate

import "testing"

func TestComposite(t *testing.T) {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")
	usrDir := NewDirectory("usr")

	binDir = binDir.Add(NewFile("vi", 10000))
	binDir = binDir.Add(NewFile("latex", 20000))

	yuki := NewDirectory("yuki")
	hanako := NewDirectory("hanako")
	tomura := NewDirectory("tomura")

	yuki = yuki.Add(NewFile("diary.html", 100))
	yuki = yuki.Add(NewFile("Composite.java", 200))

	hanako = hanako.Add(NewFile("memo.tex", 300))

	tomura = tomura.Add(NewFile("game.doc", 400))
	tomura = tomura.Add(NewFile("junk.mail", 500))

	usrDir = usrDir.Add(yuki)
	usrDir = usrDir.Add(hanako)
	usrDir = usrDir.Add(tomura)

	rootDir = rootDir.Add(binDir)
	rootDir = rootDir.Add(tmpDir)
	rootDir = rootDir.Add(usrDir)

	rootDir.PrintList()
}
