package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"text/template"
)

type Stg2MysqlMysqlDmlTableService struct {
	configs  []config.ExportConfig
	template *template.Template
}

func NewStg2MysqlMysqlDmlTableService(cfg config.Configuration) *Stg2MysqlMysqlDmlTableService {
	var temp, err = template.ParseFiles(cfg.Template)
	if err != nil {
		panic(err)
	}
	return &Stg2MysqlMysqlDmlTableService{
		configs:  cfg.Configs,
		template: temp,
	}
}

func (s *Stg2MysqlMysqlDmlTableService) PrintData(w io.Writer) error {
	for _, cfg := range s.configs {
		err := s.template.Execute(w, cfg)
		if err != nil {
			return err
		}
	}
	return nil
}
