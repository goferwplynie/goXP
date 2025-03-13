package filepicker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
	Cache       map[string][]os.DirEntry
}

type readDirMsg struct {
	Files []os.DirEntry
	err   error
}

type someMsg struct {
}

func newReadDirMsg(files []os.DirEntry, err error) tea.Msg {
	return readDirMsg{
		Files: files,
		err:   err,
	}
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
		CursorPos:   0,
		CurrentDir:  SetupPath(),
		ShowSize:    true,
		ShowMode:    true,
		ShowModTime: true,
		ShowContent: true,
		Cache:       make(map[string][]os.DirEntry),
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

func (m Model) JoinPath() (path string) {
	dirStack := m.CurrentDir
	path = ""

	for v := range dirStack.Range() {
		path = path + v + string(filepath.Separator)
	}
	return path
}

func (m *Model) ReadDir() tea.Cmd {
	return func() tea.Msg {
		path := m.JoinPath()
		for k, _ := range m.Cache {
			if k == path {
				return newReadDirMsg(m.Cache[k], nil)
			}
		}
		files, err := os.ReadDir(path)
		m.Files = files
		if err == nil {
			m.Cache[path] = files
		}
		return newReadDirMsg(files, err)

	}
}

func (m Model) Init() tea.Cmd {
	return m.ReadDir()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case readDirMsg:
		if msg.err != nil {
			panic(msg.err)
		}
		m.CursorPos = 0
		m.Files = msg.Files
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keybinds.Down):
			if m.CursorPos >= len(m.Files)-1 {
				break
			} else {
				m.CursorPos += 1
			}
		case key.Matches(msg, m.Keybinds.Up):
			if m.CursorPos < 1 {
				break
			} else {
				m.CursorPos -= 1
			}
		case key.Matches(msg, m.Keybinds.Enter):
			currentFile := m.Files[m.CursorPos]
			if !currentFile.IsDir() {
				break
			} else {
				m.CurrentDir.Append(currentFile.Name())
				return m, m.ReadDir()
			}
		case key.Matches(msg, m.Keybinds.Back):
			m.CurrentDir.Pop()
			return m, m.ReadDir()
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := ""
	s += m.JoinPath() + "\n"

	for i, v := range m.Files {
		if i == m.CursorPos {
			s += fmt.Sprintf("%s  ", m.Cursor)
		} else {
			s += fmt.Sprintf("%v. ", i)

		}
		s += fmt.Sprintf("%s", v.Name())
		if v.IsDir() {
			s += string(filepath.Separator)
		}
		s += "\n"
	}

	return s
}
