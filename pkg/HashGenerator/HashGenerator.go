package hashgenerator

const (
	localHashesFileName  = "hashes.json"
	serverHashesFileName = "hashes_server.json"
)

type hashGenerator struct {
	fileBrowser FileBrowser
	fileHandler *fileHandler
	jsonSaver   JsonSaverApi
}

func NewHashGenerator(fileHandler *fileHandler) *hashGenerator {
	return &hashGenerator{
		fileBrowser: NewFileBrowser(fileHandler),
		fileHandler: fileHandler,
		jsonSaver:   NewJsonSaver(),
	}
}

func (hg *hashGenerator) GenerateHashes(directory string) error {
	err := hg.fileBrowser.BrowseFiles(directory)
	if err != nil {
		return err
	}
	list := hg.fileHandler.GetResultList()
	var hashFileName string
	if hg.fileHandler.storageBrowser == nil {
		hashFileName = localHashesFileName
	} else {
		hashFileName = serverHashesFileName
	}
	hg.jsonSaver.SaveJson(list, hashFileName)
	return nil
}
