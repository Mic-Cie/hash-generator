package hashgenerator

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type validJsonSchema struct {
	StringField string
	IntField    int
	BoolField   bool
}

type invalidJsonSchema struct {
	FuncField func()
}

var (
	exampleValidJson = validJsonSchema{
		"String", 20, true,
	}
	exampleInvalidJson = invalidJsonSchema{
		mockVoidFunction,
	}
	errWriteFile = errors.New("mock write file error")
	errGetWd     = errors.New("mock get wd error")
)

func TestSaveJsonSuccessCase(t *testing.T) {
	jsonSaver := NewJsonSaver()
	jsonSaver.writeFile = mockWriteFileSuccess

	err := jsonSaver.SaveJson(exampleValidJson, "file.json")

	assert.Nil(t, err)
}

func TestSaveJsonFailureCase(t *testing.T) {
	tests := []struct {
		name      string
		inputJson any
		getWdFunc func() (string, error)
		expErrMsg string
	}{
		{
			name:      "should return an error, if failed to convert input into json",
			inputJson: exampleInvalidJson,
			getWdFunc: os.Getwd,
			expErrMsg: "json: unsupported type: func()",
		},
		{
			name:      "should return an error, if failed to get working dir",
			inputJson: exampleValidJson,
			getWdFunc: mockGetWdFail,
			expErrMsg: errGetWd.Error(),
		},
		{
			name:      "should return an error, if failed to write a file",
			inputJson: exampleValidJson,
			getWdFunc: os.Getwd,
			expErrMsg: errWriteFile.Error(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonSaver := NewJsonSaver()
			jsonSaver.writeFile = mockWriteFileFail
			jsonSaver.getWd = test.getWdFunc

			err := jsonSaver.SaveJson(test.inputJson, "file.json")

			require.NotNil(t, err)
			assert.Contains(t, err.Error(), test.expErrMsg)
		})
	}
}

func mockVoidFunction() {
}

func mockWriteFileFail(filename string, data []byte, perm fs.FileMode) error {
	return errWriteFile
}

func mockWriteFileSuccess(filename string, data []byte, perm fs.FileMode) error {
	return nil
}

func mockGetWdFail() (string, error) {
	return "", errGetWd
}
