package hashgenerator

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type Hasher interface {
	HashFile(file string) (string, error)
}

type FileHasher struct {
}

func NewFileHasher() FileHasher {
	return FileHasher{}
}

func (fh *FileHasher) HashFile(filePath string) (string, error) {
	hashBuffer := sha256.New()
	err := fh.fillHashBuffer(hashBuffer, filePath)
	if err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", hashBuffer.Sum(nil))
	return hash, nil
}

func (fh *FileHasher) fillHashBuffer(hashBuffer io.Writer, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(hashBuffer, f); err != nil {
		return err
	}
	return nil
}
