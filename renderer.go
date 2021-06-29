package encore

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"path"
	"strings"
)

type Renderer struct {
	entrypoints Entrypoints
	path        string
}

func NewRenderer(assetFS fs.FS, path string) *Renderer {
	entrypointFile, err := assetFS.Open(defaultEntrypointFile)
	if err != nil {
		log.Fatalf("%s file not found : %s", defaultEntrypointFile, err.Error())
	}

	entrypoints, err := readEntrypoints(entrypointFile)
	if err != nil {
		log.Fatalf("cannot read entrypoint file : %s", err.Error())
	}

	return &Renderer{
		entrypoints: entrypoints,
		path:        path,
	}
}

func (c *Renderer) FuncMap() template.FuncMap{
	return template.FuncMap{
		"encore_entry_link_tags": c.renderLinkTags,
		"encore_entry_script_tags": c.renderScriptTags,
	}
}

func (c *Renderer) renderLinkTags(entrypoint string) template.HTML {
	var links []string
	for _, fileName := range c.entrypoints[entrypoint].CSS {
		links = append(links, fmt.Sprintf(`<link rel="stylesheet" href="%s">`, path.Join(c.path, fileName)))
	}

	return template.HTML(strings.Join(links, "\n"))
}

func (c *Renderer) renderScriptTags(entrypoint string) template.HTML {
	var scripts []string
	for _, fileName := range c.entrypoints[entrypoint].JS {
		scripts = append(scripts, fmt.Sprintf(`<script src="%s" defer></script>`, path.Join(c.path, fileName)))
	}

	return template.HTML(strings.Join(scripts, "\n"))
}

