package filepicker

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/goferwplynie/goXP/internal/ds/linkedlist"
	"github.com/goferwplynie/goXP/internal/ds/stack"
)

type KeyBinds struct {
	Up         key.Binding
	Down       key.Binding
	Back       key.Binding
	CmdMode    key.Binding
	SelectMode key.Binding
	SelectOne  key.Binding
	Enter      key.Binding
	Delete     key.Binding
	Add        key.Binding
	AddDir     key.Binding
	Rename     key.Binding
	Undo       key.Binding
	Redo       key.Binding
}

type Model struct {
	Files       []os.DirEntry
	Keybinds    KeyBinds
	Cursor      string
	CursorPos   int
	Selected    []int
	CurrentDir  stack.Stack[string]
	ShowSize    bool
	ShowMode    bool
	ShowModTime bool
	ShowContent bool
}

func DefaultKeyBinds() KeyBinds {
	return KeyBinds{
		Up:         key.NewBinding(key.WithKeys("k", "up")),
		Down:       key.NewBinding(key.WithKeys("j", "down")),
		Back:       key.NewBinding(key.WithKeys("h", "left")),
		CmdMode:    key.NewBinding(key.WithKeys(":", "/")),
		SelectMode: key.NewBinding(key.WithKeys("V")),
		SelectOne:  key.NewBinding(key.WithKeys("v")),
		Enter:      key.NewBinding(key.WithKeys("l", "enter")),
		Delete:     key.NewBinding(key.WithKeys("d", "backspace")),
		Add:        key.NewBinding(key.WithKeys("a", "n")),
		AddDir:     key.NewBinding(key.WithKeys("A", "N")),
		Rename:     key.NewBinding(key.WithKeys("r")),
		Undo:       key.NewBinding(key.WithKeys("u", "ctrl+z")),
		Redo:       key.NewBinding(key.WithKeys("ctrl+r", "ctrl+y", "ctrl+Z")),
	}
}

func New() Model {
	return Model{
		Files:       nil,
		Keybinds:    DefaultKeyBinds(),
		Cursor:      ">",
		CurrentDir:  SetupPath(),
		ShowSize:    true,
		ShowMode:    true,
		ShowModTime: true,
		ShowContent: true,
	}
}

func SetupPath() *linkedlist.LinkedList[string] {
	ll := linkedlist.NewLinkedList[string]()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(path, string(filepath.Separator))
	for _, v := range dirs {
		ll.Append(v)
	}
	return ll
}
