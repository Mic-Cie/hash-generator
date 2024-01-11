package hashgenerator

import (
	"strings"
)

// excludedFilesNames are names that should be excluded from handling
var excludedFileNames = []string{localHashesFileName, serverHashesFileName}

type FileHandler interface {
	Handle(path string) error
}

type FileInfo struct {
	FilePath string
	Hash     string
	NodeHash string
}

type fileHandler struct {
	resultList     []FileInfo
	hasher         Hasher
	storageBrowser StorageBrowser
}

func NewFileHandler(storageBrowser StorageBrowser) fileHandler {
	hasher := NewFileHasher()
	return fileHandler{
		hasher:         &hasher,
		storageBrowser: storageBrowser,
	}
}

func (fh *fileHandler) Handle(path string) error {
	if FileNameContainsExcludedName(path) {
		return nil
	}
	hash, err := fh.hasher.HashFile(path)
	if err != nil {
		return err
	}

	var nodeHash string
	if fh.storageBrowser != nil {
		nodeHash, err = fh.storageBrowser.GetObjectNode(path)
		if err != nil {
			return err
		}
	}

	fileInfo := FileInfo{FilePath: path, Hash: hash, NodeHash: nodeHash}
	fh.resultList = append(fh.resultList, fileInfo)
	return nil
}

func (fh *fileHandler) GetResultList() []FileInfo {
	return fh.resultList
}

func FileNameContainsExcludedName(path string) bool {
	for _, ex := range excludedFileNames {
		if strings.Contains(path, ex) {
			return true
		}
	}
	return false
}
