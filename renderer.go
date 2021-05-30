package encore

import (
	"fmt"
	"html/template"
	"path"
	"strings"
)

type Renderer struct {
	entrypoints Entrypoints
	path        string
}

func NewRenderer(entrypoints Entrypoints, path string) *Renderer {
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

