package templates

import (
	"errors"
	"html/template"
	"platform/config"
	"sync"
)

var once sync.Once

func LoadTemplates(cfg config.Configuration) error {
	path, found := cfg.GetString("templates:path")
	if !found {
		return errors.New("cannot find templates path settings")
	}
	var err error
	reload := cfg.GetBoolDefault("templates:reload", false)
	once.Do(func() {
		doLoad := func() *template.Template {
			t := template.New("htmlTemplates")
			t.Funcs(map[string]any{
				"body":   func() string { return "" },
				"layout": func() string { return "" },
			})
			t, err = t.ParseGlob(path)
			return t
		}
		if reload {
			getTemplates = doLoad
			return
		}
		templates := doLoad()
		getTemplates = func() *template.Template {
			t, _ := templates.Clone()
			return t
		}
	})
	return err
}
