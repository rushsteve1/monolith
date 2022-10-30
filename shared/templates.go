package shared

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type TemplateHelper struct {
	inner map[string]*template.Template
}

func LoadTemplates(fs embed.FS, dir string, layout string) (*TemplateHelper, error) {
	th := TemplateHelper{
		inner: make(map[string]*template.Template),
	}

	files, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Name() == layout {
			continue
		}

		base := template.New("")
		base.Funcs(template.FuncMap{
			"safe": func(s string) template.HTML {
				return template.HTML(s)
			}})

		log.Trace("Parsing template ", file.Name())

		t, err := base.ParseFS(fs, filepath.Join(dir, file.Name()), filepath.Join(dir, layout))
		if err != nil {
			return nil, fmt.Errorf("Failed to load templates: %v", err)
		}

		th.inner[file.Name()] = t.Lookup(file.Name())
	}

	return &th, nil
}

func (th *TemplateHelper) Render(w io.Writer, name string, data any) error {
	return th.inner[name].Execute(w, data)
}
