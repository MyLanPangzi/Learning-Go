package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"os"
	"text/template"
)

type Stg2MysqlHiveDdlTableService struct {
	configs     []config.ExportConfig
	endpoint    string
	dbService   DbService
	ddlTemplate *template.Template
}

func NewStg2MysqlHiveDdlTableService(cfg config.Configuration, dbService DbService) *Stg2MysqlHiveDdlTableService {
	bytes, err := os.ReadFile(cfg.Template)
	if err != nil {
		panic(err)
	}
	temp, err := template.New("ddl").Parse(string(bytes))
	if err != nil {
		panic(err)
	}
	return &Stg2MysqlHiveDdlTableService{
		configs:     cfg.Configs,
		endpoint:    cfg.Endpoint,
		dbService:   dbService,
		ddlTemplate: temp,
	}
}

func (s *Stg2MysqlHiveDdlTableService) PrintData(writer io.Writer) error {
	return s.ddlTemplate.Execute(writer, s.GetData())
}
func (s *Stg2MysqlHiveDdlTableService) GetData() []TableData {
	var out []TableData
	for _, cfg := range s.configs {
		out = append(out, s.makeData(cfg)...)
	}
	return out
}

func (s *Stg2MysqlHiveDdlTableService) makeData(cfg config.ExportConfig) []TableData {
	var out []TableData
	db, err := s.dbService.OpenDb(cfg, "information_schema", s.endpoint)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("set group_concat_max_len =1024*1024*1024 ")
	if err != nil {
		panic(err)
	}
	for _, t := range cfg.Tables {

		row := db.QueryRow(
			"select group_concat(concat('`',"+"COLUMN_NAME,'`', "+`'\t', case
                                                  when DATA_TYPE like '%int' then DATA_TYPE
                                                  when DATA_TYPE like '%char' then COLUMN_TYPE
                                                  when COLUMN_TYPE = 'timestamp' then 'TIMESTAMP_LTZ(3)'
                                                  when DATA_TYPE = 'timestamp'
                                                      then replace(COLUMN_TYPE, 'timestamp', 'TIMESTAMP_LTZ')
                                                  when DATA_TYPE = 'bit' then 'int'
                                                  when DATA_TYPE = 'json' then 'string'
                                                  when DATA_TYPE = 'text' then 'string'
                                                  else COLUMN_TYPE
							end) separator ',\n\t') columns
					from COLUMNS
					where TABLE_SCHEMA = ?
					and TABLE_NAME = ?
					order by ORDINAL_POSITION
		`,
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
