package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"text/template"
)

type Stg2HiveDdlTableService struct {
	configs  []config.ExportConfig
	template *template.Template
}

func NewStg2HiveTableService(cfg config.Configuration) *Stg2HiveDdlTableService {
	template, err := template.ParseFiles(cfg.Template)
	if err != nil {
		panic(err)
	}
	return &Stg2HiveDdlTableService{
		configs:  cfg.Configs,
		template: template,
	}
}

func (s *Stg2HiveDdlTableService) PrintData(w io.Writer) error {
	type Data struct {
		Db    string
		Table string
	}
	data := make([]Data, 0, 10)
	for _, cfg := range s.configs {
		for _, t := range cfg.Tables {
			data = append(data, Data{cfg.Db, t.Table})
		}
	}
	return s.template.Execute(w, data)
}
