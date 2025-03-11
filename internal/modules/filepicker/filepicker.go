package filepicker

import (
	"os"
)

type Model struct {
	Files       []os.DirEntry
	CurrentDir  string
	ShowSize    bool
	ShowMode    bool
	ShowModTime bool
	ShowContent bool
}

func main() {

}
