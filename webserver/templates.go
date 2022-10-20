package webserver

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*.html
var templatesFS embed.FS
var templates *template.Template

func loadTemplates() error {
	var err error
	templates, err = template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return fmt.Errorf("Failed to load templates: %v", err)
	}
	return nil
}
