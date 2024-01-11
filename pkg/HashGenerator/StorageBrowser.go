package hashgenerator

type StorageBrowser interface {
	GetObjectNode(file string) (string, error)
}

type MegaBrowser struct {
}

func NewMegaBrowser() (*MegaBrowser, error) {
	return &MegaBrowser{}, nil
}

func (b *MegaBrowser) GetObjectNode(file string) (string, error) {
	return "", nil
}
