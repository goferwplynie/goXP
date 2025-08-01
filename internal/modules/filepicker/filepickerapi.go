package filepicker

import (
	"io"
	"os"
	"regexp"
)

type FilePickerApi interface {
	GetFiles() []os.DirEntry
	GetSelected() []os.DirEntry
	Search(regex string) []os.DirEntry
	GetCurrentDir() string
	GetCache() map[string][]os.DirEntry
	RenameFile(file string, newName string) bool
	DeleteFile(file string) bool
	OverwriteFiles(files []os.DirEntry)
	SetDirectory(dir string)
}

func (m *Model) SetDirectory(dir string) {

}

func (m *Model) GetFiles() []os.DirEntry {
	return m.Files
}

func (m *Model) GetSelected() []os.DirEntry {
	return m.Selected
}

func (m *Model) Search(regex string) []os.DirEntry {
	files := m.GetFiles()

	var matching []os.DirEntry

	for _, v := range files {
		match, err := regexp.MatchString(regex, v.Name())
		if err != nil {
			panic(err)
		}
		if match {
			matching = append(matching, v)
		}

	}
	return matching
}

func (m *Model) GetCurrentDir() string {
	return m.JoinPath()
}

func (m *Model) GetCache() map[string][]os.DirEntry {
	return m.Cache
}

// users can make their own sort or filter files
func (m *Model) OverwriteFiles(files []os.DirEntry) {
	m.Files = files
}

func (m *Model) DeleteFile(file string) bool {
	currentPath := m.GetCurrentDir()
	filePath := currentPath + file
	err := os.Remove(filePath)
	if err != nil {
		return false
	}
	return true
}

func (m *Model) RenameFile(file string, newName string) bool {
	path := m.GetCurrentDir() + file
	newPath := m.GetCurrentDir() + newName
	err := os.Rename(path, newPath)
	if err != nil {
		return false
	}

	return true
}

func MoveFile(src, dst string) bool {
	err := os.Rename(src, dst)
	// trying with os.rename (faster)
	if err == nil {
		return true
	}

	//if failed trying full moving
	sourceFile, err := os.Open(src)
	if err != nil {
		return false
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return false
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return false
	}

	err = os.Remove(src)
	if err != nil {
		return false
	}

	return true
}
