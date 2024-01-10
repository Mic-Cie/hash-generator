package hashgenerator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errHash = errors.New("mock hash error")

type mockHasher struct {
	err error
}

func TestFileHandler(t *testing.T) {
	tests := []struct {
		name   string
		paths  []string
		expRes []FileInfo
		expErr error
	}{
		{
			name:  "should generate a list, if given some filepaths",
			paths: []string{"file1", "file2"},
			expRes: []FileInfo{
				{"file1", "file1", ""},
				{"file2", "file2", ""},
			},
			expErr: nil,
		},
		{
			name:  "should generate a list, if given some filepaths, and omit 'uploader.exe'",
			paths: []string{"file1", "file2", excludedFileNames[0]},
			expRes: []FileInfo{
				{"file1", "file1", ""},
				{"file2", "file2", ""},
			},
			expErr: nil,
		},
		{
			name:   "should leave the list empty, if no filepaths were given",
			paths:  nil,
			expRes: nil,
			expErr: nil,
		},
		{
			name:   "should return an error, if failed to generate a hash",
			paths:  []string{"invalid path"},
			expRes: nil,
			expErr: errHash,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewFileHandler()
			mockHasher := &mockHasher{test.expErr}
			handler.hasher = mockHasher
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
