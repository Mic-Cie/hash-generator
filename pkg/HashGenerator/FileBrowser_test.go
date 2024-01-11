package hashgenerator

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDir = "testDir"

var errHandleFile = fmt.Errorf("mock file handle error")

type fileHandlerMock struct {
	fileList []string
	err      error
}

func TestFileBrowserSuccessCase(t *testing.T) {
	expList := []string{
		filepath.Join(testDir, "file1.txt"),
		filepath.Join(testDir, "file2.txt"),
		filepath.Join(testDir, "subDir", "file3.txt"),
	}
	handler := fileHandlerMock{
		err: nil,
	}

	fl := NewFileBrowser(&handler)
	err := fl.BrowseFiles(testDir)

	assert.Equal(t, expList, handler.fileList)
	assert.Nil(t, err)
}

func TestFileBrowserFailCase(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		handleErr error
		expErrMsg string
	}{
		{
			name:      "should return an error, if failed to handle a file",
			path:      "testDir",
			handleErr: errHandleFile,
			expErrMsg: errHandleFile.Error(),
		},
		{
			name:      "should return an error, if failed to browse directory",
			path:      "dirDoesntExist",
			handleErr: nil,
			expErrMsg: "CreateFile dirDoesntExist:",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := fileHandlerMock{
				err: test.handleErr,
			}

			fl := NewFileBrowser(&handler)
			err := fl.BrowseFiles(test.path)
			assert.Nil(t, handler.fileList)
			require.NotNil(t, err)
			assert.Contains(t, err.Error(), test.expErrMsg)
		})
	}
}

func (fh *fileHandlerMock) Handle(path string) error {
	if fh.err != nil {
		return fh.err
	}
	fh.fileList = append(fh.fileList, path)
	return nil
}
