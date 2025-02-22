package filepicker

import (
	"errors"
	"log"
	"os"

	"github.com/goferwplynie/goXP/config"
	"github.com/goferwplynie/goXP/internal/models"
	"path/filepath"
)

type FilePicker struct {
	Files             []models.File
	CurrentFolderName string
	CurrentPath       string
	Cursor            string
	Selected          int
}

var conf = config.New()

func New() *FilePicker{
	return &FilePicker{
		Files: []models.File{},


	}
}

func (fp *FilePicker) FindFiles() error {
	files, err := os.ReadDir(fp.CurrentPath)

	if err != nil {
		return errors.New("Unable to read this directory")
	}

	if len(files) > conf.MaxFilesToLoadAtStart {
		for i := 0; i <= conf.MaxFilesToLoadAtStart; i++ {
			file := files[i]
			info, err := file.Info()
			if err != nil {
				log.Printf("%v: %v", file.Name(), err)
			}

			fileObj := models.File{
				Name:    info.Name(),
				Size:    info.Size(),
				Mode:    uint32(info.Mode()),
				ModTime: info.ModTime(),
				IsDir:   info.IsDir(),
			}

			fp.Files = append(fp.Files, fileObj)
		}
		files = files[conf.MaxFilesToLoadAtStart:]

		go func() {
			for _, file := range files {
				info, err := file.Info()
				if err != nil {
					log.Printf("%v: %v", file.Name(), err)
				}

				fileObj := models.File{
					Name:    info.Name(),
					Size:    info.Size(),
					Mode:    uint32(info.Mode()),
					ModTime: info.ModTime(),
					IsDir:   info.IsDir(),
				}

				fp.Files = append(fp.Files, fileObj)
			}
		}()
	}

	return nil
}
