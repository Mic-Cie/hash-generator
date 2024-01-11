package hashgenerator

import (
	"os"
	"path/filepath"
)

type FileBrowser struct {
	fileHandler FileHandler
}

func NewFileBrowser(FileHandler FileHandler) FileBrowser {
	return FileBrowser{
		fileHandler: FileHandler,
	}
}

func (fl *FileBrowser) BrowseFiles(dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				handleErr := fl.fileHandler.Handle(path)
				if handleErr != nil {
					return handleErr
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
