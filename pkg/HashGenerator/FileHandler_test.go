package hashgenerator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errHash = errors.New("mock hash error")
var errGetNode = errors.New("mock get node error")

type mockHasher struct {
	err error
}

type mockBrowser struct {
	err error
}

func TestFileHandler(t *testing.T) {
	tests := []struct {
		name       string
		paths      []string
		hashErr    error
		getNodeErr error
		expRes     []FileInfo
		expErr     error
	}{
		{
			name:       "should generate a list, if given some filepaths",
			paths:      []string{"file1", "file2"},
			hashErr:    nil,
			getNodeErr: nil,
			expRes: []FileInfo{
				{"file1", "file1", ""},
				{"file2", "file2", ""},
			},
			expErr: nil,
		},
		{
			name:       "should generate a list, if given some filepaths, and omit 'hashes.json'",
			paths:      []string{"file1", "file2", excludedFileNames[0]},
			hashErr:    nil,
			getNodeErr: nil,
			expRes: []FileInfo{
				{"file1", "file1", ""},
				{"file2", "file2", ""},
			},
			expErr: nil,
		},
		{
			name:       "should leave the list empty, if no filepaths were given",
			paths:      nil,
			hashErr:    nil,
			getNodeErr: nil,
			expRes:     nil,
			expErr:     nil,
		},
		{
			name:       "should return an error, if failed to generate a hash",
			paths:      []string{"invalid path"},
			hashErr:    errHash,
			getNodeErr: nil,
			expRes:     nil,
			expErr:     errHash,
		},
		{
			name:       "should return an error, if failed to get object node",
			paths:      []string{"file1", "file2"},
			hashErr:    nil,
			getNodeErr: errGetNode,
			expRes:     nil,
			expErr:     errGetNode,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewFileHandler()
			mockHasher := &mockHasher{test.hashErr}
			mockBrowser := &mockBrowser{test.getNodeErr}
			handler.hasher = mockHasher
			handler.storageBrowser = mockBrowser
			for _, path := range test.paths {
				err := handler.Handle(path)
				require.Equal(t, test.expErr, err)
			}
			assert.Equal(t, test.expRes, handler.GetResultList())
		})
	}
}

func (m *mockHasher) HashFile(file string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return file, nil
}

func (m *mockBrowser) GetObjectNode(file string) (string, error) {
	return "", m.err
}
