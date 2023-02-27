package services

import (
	"platform/config"
	"platform/logging"
	"platform/templates"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() config.Configuration {
		cfg, err := config.Load("config.json")
		if err != nil {
			panic(err)
		}
		return cfg
	})
	if err != nil {
		panic(err)
	}
	err = AddSingleton(func(cfg config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(cfg)
	})
	if err != nil {
		panic(err)
	}
	err = AddSingleton(func(cfg config.Configuration) templates.TemplateExecutor {
		err := templates.LoadTemplates(cfg)
		if err != nil {
			panic(err)
		}
		return &templates.LayoutTemplateProcessor{}
	})
	if err != nil {
		panic(err)
	}
}
