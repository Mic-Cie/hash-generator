package hashgenerator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFilePath = "testDir/file1.txt"

var errCopyBuffer = fmt.Errorf("mock error")

type mockWriter struct {
}

func TestShouldReturnHashIfSuccessfullyLoadedAFile(t *testing.T) {
	expHashLength := 64

	fh := NewFileHasher()
	hash, err := fh.HashFile(testFilePath)
	assert.Equal(t, expHashLength, len(hash))
	assert.Nil(t, err)
}

func TestShouldFailIfFailedOpeningAFile(t *testing.T) {
	fh := NewFileHasher()
	hash, err := fh.HashFile("")
	assert.Equal(t, "", hash)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "open :")
}

func TestShouldFailIfCouldNotCopyBuffer(t *testing.T) {
	fh := NewFileHasher()
	mockBuffer := mockWriter{}
	err := fh.fillHashBuffer(&mockBuffer, testFilePath)
	require.NotNil(t, err)
	assert.Equal(t, errCopyBuffer, err)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return 0, errCopyBuffer
}
