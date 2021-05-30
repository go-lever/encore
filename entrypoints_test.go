package encore_test

import (
	"testing"

	"github.com/Flaque/filet"
	"github.com/seblegall/encore"
	"github.com/stretchr/testify/assert"
)

func TestReadEntrypoints(t *testing.T) {
	defer filet.CleanUp(t)

	// Creates a temporary file with string "some content"
	f := filet.TmpFile(t, "", `
{
  "entrypoints": {
    "app": {
      "js": [
        "/runtime.c7f5ad70.js",
        "/app.3be12169.js"
      ],
      "css": [
        "/app.7da538eb.css"
      ]
    },
    "page1": {
      "js": [
        "/runtime.c7f5ad70.js",
        "/page1.4eb517cf.js"
      ],
      "css": [
        "/page1.f49d9db2.css"
      ]
    }
  }
}`)
	entry, err  := encore.ReadEntrypoints(f.Name())
	assert.Nil(t, err)
	_, ok := entry["app"]
	assert.True(t, ok, "entrypoint app not found")
	page1, ok := entry["page1"]
	assert.True(t, ok, "page1 app not found")
	assert.Equal(t, 2, len(page1.JS))
}