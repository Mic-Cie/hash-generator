package hashgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockJsonSaver struct {
}

func TestHashGeneratorSuccessCase(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		storageBrowser StorageBrowser
	}{
		{
			name:           "should generate local hashes, if the directory is valid",
			path:           "testDir",
			storageBrowser: nil,
		},
		{
			name:           "should generate server hashes, if the directory is valid and is using the storage",
			path:           "testDir",
			storageBrowser: &MegaBrowser{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fileHandler := NewFileHandler(test.storageBrowser)
			hg := NewHashGenerator(&fileHandler)
			jsonSaver := mockJsonSaver{}
			hg.jsonSaver = &jsonSaver

			require.NotNil(t, hg)
			err := hg.GenerateHashes(test.path)
			assert.Nil(t, err)
		})
	}
}

func TestHashGeneratorFailureCase(t *testing.T) {
	fileHandler := NewFileHandler(nil)
	hg := NewHashGenerator(&fileHandler)
	jsonSaver := mockJsonSaver{}
	hg.jsonSaver = &jsonSaver

	err := hg.GenerateHashes("invalid")

	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid:")
}

func (m *mockJsonSaver) SaveJson(v any, fileName string) error {
	return nil
}
