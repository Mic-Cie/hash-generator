package hashgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetObjectNode(t *testing.T) {
	storageBrowser := StorageBrowser{}
	node, err := storageBrowser.GetObjectNode("")
	assert.Equal(t, "", node)
	assert.Nil(t, err)
}
