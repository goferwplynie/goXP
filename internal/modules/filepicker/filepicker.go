package filepicker

import (
	"os"

	"github.com/charmbracelet/bubbles/key"
)

type KeyBinds struct {
	Up      key.Binding
	Down    key.Binding
	Back    key.Binding
	CmdMode key.Binding
	Select  key.Binding
	Enter   key.Binding
	Delete  key.Binding
	Add     key.Binding
	Rename  key.Binding
}

type Model struct {
	Files       []os.DirEntry
	Keybinds    KeyBinds
	Cursor      string
	CurrentDir  string
	ShowSize    bool
	ShowMode    bool
	ShowModTime bool
	ShowContent bool
}

func DefaultKeyBinds() KeyBinds {
	return KeyBinds{
		Up: key.NewBinding(key.WithKeys("k", "up")),
	}
}
