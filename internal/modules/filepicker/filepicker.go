package filepicker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goferwplynie/goXP/config"
	"github.com/goferwplynie/goXP/internal/ds/linkedlist"
	"github.com/goferwplynie/goXP/internal/ds/stack"
	"github.com/goferwplynie/goXP/internal/styles"
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
	Selected    []os.DirEntry
	CurrentDir  stack.Stack[string]
	ShowSize    bool
	ShowMode    bool
	ShowModTime bool
	ShowContent bool
	Cache       map[string][]os.DirEntry
	Styles      FilePickerStyle
}

type FilePickerStyle struct {
	CurrentFile  lipgloss.Style
	DefaultFile  lipgloss.Style
	Folder       lipgloss.Style
	CurrentPath  lipgloss.Style
	ModeStyle    lipgloss.Style
	ModTimeStyle lipgloss.Style
	SizeStyle    lipgloss.Style
	Selected     lipgloss.Style
	CursorStyle  lipgloss.Style
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

func CustomStyle(fpStyles config.FilePickerStyles) FilePickerStyle {
	return FilePickerStyle{
		CurrentFile:  styles.BuildStyle(fpStyles.CurrentFile),
		DefaultFile:  styles.BuildStyle(fpStyles.DefaultFile),
		Folder:       styles.BuildStyle(fpStyles.Folder),
		CurrentPath:  styles.BuildStyle(fpStyles.CurrentPath),
		ModeStyle:    styles.BuildStyle(fpStyles.ModeStyle),
		ModTimeStyle: styles.BuildStyle(fpStyles.ModTimeStyle),
		SizeStyle:    styles.BuildStyle(fpStyles.SizeStyle),
		Selected:     styles.BuildStyle(fpStyles.Selected),
		CursorStyle:  styles.BuildStyle(fpStyles.CursorStyle),
	}

}

func CustomKeybinds(c config.FilePickerKeybinds) KeyBinds {
	return KeyBinds{
		Up:         key.NewBinding(key.WithKeys(c.Up...)),
		Down:       key.NewBinding(key.WithKeys(c.Down...)),
		Back:       key.NewBinding(key.WithKeys(c.Back...)),
		CmdMode:    key.NewBinding(key.WithKeys(":", "/")),
		SelectMode: key.NewBinding(key.WithKeys(c.SelectMode...)),
		SelectOne:  key.NewBinding(key.WithKeys(c.SelectOne...)),
		Enter:      key.NewBinding(key.WithKeys(c.Enter...)),
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
		CursorPos:   0,
		CurrentDir:  SetupPath(),
		ShowSize:    true,
		ShowMode:    true,
		ShowModTime: false,
		ShowContent: true,
		Cache:       make(map[string][]os.DirEntry),
	}
}

func SetupPath() stack.Stack[string] {
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
		_, exists := m.Cache[path]
		if exists {
			return newReadDirMsg(m.Cache[path], nil)
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
	s += m.Styles.CurrentPath.Render(m.JoinPath()) + "\n"

	for i, v := range m.Files {
		var filename string
		row := ""
		info, err := v.Info()
		if err != nil {
			break
		}
		if i == m.CursorPos {
			row += fmt.Sprintf("%s  ", m.Cursor)
		} else {
			row += fmt.Sprintf("%-4v ", fmt.Sprintf("%v.", i))
		}
		if v.IsDir() {
			filename = v.Name() + string(filepath.Separator)
		} else {
			filename = v.Name()
		}
		row += fmt.Sprintf("%-20s", filename)

		if m.ShowMode {
			row += m.Styles.ModeStyle.Render(fmt.Sprintf(" %-10v ", info.Mode()))
		}
		if m.ShowModTime {
			row += m.Styles.ModTimeStyle.Render(fmt.Sprintf(" %v ", info.ModTime()))
		}
		if m.ShowSize {
			row += m.Styles.SizeStyle.Render(fmt.Sprintf(" %v ", info.Size()))
		}

		if i != m.CursorPos {
			row = m.Styles.DefaultFile.Render(row)
		}
		if v.IsDir() {
			row = m.Styles.Folder.Render(row)
		}
		if i == m.CursorPos {
			row = m.Styles.CurrentFile.Render(row)
		}
		s += fmt.Sprintf("%s\n", row)
	}

	return s
}
