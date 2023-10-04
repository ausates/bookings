package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the site-wide application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
