package lighthouse_test

import (
	"embed"
	"io/ioutil"
	"path"
	"path/filepath"
	"testing"

	"github.com/Melvillian/sedge/internal/images/validator-import/lighthouse"
	"github.com/stretchr/testify/assert"
)

//go:embed context
var lighthouseContext embed.FS

func TestInitContext(t *testing.T) {
	entries, err := lighthouseContext.ReadDir("context")
	assert.NoError(t, err)
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			contextDir, err := lighthouse.InitContext()
			assert.NoError(t, err)
			data, err := ioutil.ReadFile(filepath.Join(contextDir, entry.Name()))
			assert.NoError(t, err)
			expectedData, err := lighthouseContext.ReadFile(path.Join("context", entry.Name()))
			assert.NoError(t, err)
			assert.Equal(t, expectedData, data)
		})
	}
}
