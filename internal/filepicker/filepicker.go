package filepicker

import (
	"errors"
	"os"

	"github.com/goferwplynie/goXP/internal/models"
)

type Model struct {
	Files             []models.File
	CurrentFolderName string
	CurrentPath       string
	Cursor            string
	Selected          int
}

func (m *Model) FindFiles() error {
	files, err := os.ReadDir(m.CurrentPath)

	if err != nil {
		return errors.New("Unable to read this directory")
	}
}
