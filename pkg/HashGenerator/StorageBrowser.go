package hashgenerator

type StorageBrowser interface {
	GetObjectNode(file string) (string, error)
}

type MegaBrowser struct {
}

func (b *MegaBrowser) GetObjectNode(file string) (string, error) {
	return "", nil
}
