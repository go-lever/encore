package encore_test

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/seblegall/encore"
	"github.com/stretchr/testify/assert"
)

func TestRenderLinkTags(t *testing.T) {
	enc := encore.New(encore.Entrypoints{
		"app" : encore.Entrypoint{
			CSS: []string{"/app.css", "/test.css"},
		},
	}, "/assets/")

	funcMap := enc.FuncMap()
	_, ok := funcMap["encore_entry_link_tags"]
	assert.True(t, ok, "encore_entry_link_tags is not part of the FuncMap")

	const html = `
{{encore_entry_link_tags .}}
`
	tmpl, err := template.New("html").Funcs(funcMap).Parse(html)
	assert.Nil(t, err)

	var tags bytes.Buffer
	err = tmpl.Execute(&tags, "app")
	assert.Nil(t, err)
	assert.Equal(t, `
<link rel="stylesheet" href="/assets/app.css">
<link rel="stylesheet" href="/assets/test.css">
`, tags.String())
}


func TestRenderScriptTags(t *testing.T) {
	enc := encore.New(encore.Entrypoints{
		"app" : encore.Entrypoint{
			JS: []string{"/app.js", "/page1.js"},
		},
	}, "/assets/")

	funcMap := enc.FuncMap()
	_, ok := funcMap["encore_entry_script_tags"]
	assert.True(t, ok, "encore_entry_script_tags is not part of the FuncMap")

	const html = `
{{encore_entry_script_tags .}}
`
	tmpl, err := template.New("html").Funcs(funcMap).Parse(html)
	assert.Nil(t, err)

	var tags bytes.Buffer
	err = tmpl.Execute(&tags, "app")
	assert.Nil(t, err)
	assert.Equal(t, `
<script src="/assets/app.js" defer></script>
<script src="/assets/page1.js" defer></script>
`, tags.String())
}
