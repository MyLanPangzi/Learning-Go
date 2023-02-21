package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"text/template"
)

type Stg2MysqlMysqlDdlTableService struct {
	configs     []config.ExportConfig
	endpoint    string
	dbService   DbService
	ddlTemplate *template.Template
}

func (s *Stg2MysqlMysqlDdlTableService) PrintData(writer io.Writer) error {
	return s.ddlTemplate.Execute(writer, s.GetData())
}

func NewStg2MysqlMysqlDdlTableService(configs []config.ExportConfig, endpoint string, dbService DbService) *Stg2MysqlMysqlDdlTableService {
	const text = `
{{- range . -}}
    CREATE TABLE IF NOT EXISTS {{.Db}}.{{.Table}} (
        {{.Columns}},
)
{{ end }}

`
	var temp, err = template.New("ddl").Parse(text)
	if err != nil {
		panic(err)
	}
	return &Stg2MysqlMysqlDdlTableService{
		configs:     configs,
		endpoint:    endpoint,
		dbService:   dbService,
		ddlTemplate: temp,
	}
}

func (s *Stg2MysqlMysqlDdlTableService) GetData() []TableData {
	var out []TableData
	for _, cfg := range s.configs {
		out = append(out, s.makeData(cfg)...)
	}
	return out
}

func (s *Stg2MysqlMysqlDdlTableService) makeData(cfg config.ExportConfig) []TableData {
	var out []TableData
	db, err := s.dbService.OpenDb(cfg, s.endpoint)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("set group_concat_max_len =1024*1024*1024 ")
	if err != nil {
		panic(err)
	}
	for _, t := range cfg.Tables {
		const query = `
select group_concat(concat(COLUMN_NAME,'\t' ,COLUMN_TYPE) separator ',\n') columns
from COLUMNS
where TABLE_SCHEMA = ?
  and TABLE_NAME = ?
order by ORDINAL_POSITION
`

		row := db.QueryRow(
			query,
			cfg.Db,
			t.Regexp,
		)
		var columns string
		err = row.Scan(&columns)
		if err != nil {
			panic(err)
		}
		out = append(out, TableData{cfg.Db, t.Table, columns})
	}
	return out
}
