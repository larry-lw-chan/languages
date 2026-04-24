package config

import "html/template"

// Holds Application Configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
