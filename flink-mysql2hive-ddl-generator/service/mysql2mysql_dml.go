package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"text/template"
)

type Stg2MysqlMysqlDmlTableService struct {
	configs   []config.ExportConfig
	endpoint  string
	dbService DbService
	template  *template.Template
}

func NewStg2MysqlMysqlDmlTableService(cfg config.Configuration, dbService DbService) *Stg2MysqlMysqlDmlTableService {
	const text = `
{{- $db := .Db -}}
{{- $dest := .DestDb -}}
{{ range .Tables -}}

  INSERT INTO mysql.{{ $dest }}.{{.Table}}
  /*+ OPTIONS(
  'sink.buffer-flush.max-rows' = '1000',
  'sink.buffer-flush.interval' = '3s'
  ) */
  SELECT *
  FROM hive.kafka.stg_{{$db}}_{{.Table}}_realtime_debezium_json
  /*+ OPTIONS(
  'properties.group.id' = 'stg_{{$db}}_{{.Table}}_realtime_2_mysql'
  ) */
  ;
{{- end }}

`
	var temp, err = template.New("dml").Parse(text)
	if err != nil {
		panic(err)
	}
	return &Stg2MysqlMysqlDmlTableService{
		configs:   cfg.Configs,
		endpoint:  cfg.Endpoint,
		dbService: dbService,
		template:  temp,
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
