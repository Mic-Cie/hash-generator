package hashgenerator

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

type JsonSaverApi interface {
	SaveJson(v any, fileName string) error
}

type JsonSaver struct {
	writeFile func(filename string, data []byte, perm fs.FileMode) error
	getWd     func() (string, error)
}

func NewJsonSaver() *JsonSaver {
	return &JsonSaver{
		writeFile: os.WriteFile,
		getWd:     os.Getwd,
	}
}

func (s *JsonSaver) SaveJson(v any, fileName string) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	path, err := s.getWd()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(path, fileName)
	err = s.writeFile(fullPath, body, 0666)
	if err != nil {
		return err
	}
	return nil
}
