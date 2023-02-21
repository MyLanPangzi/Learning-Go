package service

import (
	"flink-mysql-ddl-generator/config"
	"io"
	"text/template"
)

type Stg2MysqlHiveDdlTableService struct {
	configs     []config.ExportConfig
	endpoint    string
	dbService   DbService
	ddlTemplate *template.Template
}

func NewStg2MysqlHiveDdlTableService(configs []config.ExportConfig, endpoint string, dbService DbService) *Stg2MysqlHiveDdlTableService {
	const text = `
{{- range . -}}
    CREATE TABLE IF NOT EXISTS hive.kafka.stg_{{.Db}}_{{.Table}}_cdc_realtime_debezium_json (
        {{.Columns}},
        metadata_schema_name STRING METADATA FROM 'value.source.database',
        metadata_table_name STRING METADATA FROM 'value.source.table',
        metadata_timestampTIMESTAMP_LTZ(3) METADATA FROM 'value.ingestion-timestamp',
        kafka_topic         STRING METADATA FROM 'topic',
        kafka_partition     INT METADATA FROM 'partition',
        kafka_offset        BIGINT METADATA FROM 'offset',
        kafka_timestamp     TIMESTAMP(3) METADATA FROM 'timestamp',
        kafka_timestamp_typeSTRING METADATA FROM 'timestamp-type'
    )WITH (
        'connector' = 'kafka',
        'topic' = 'stg_{{.Db}}_{{.Table}}_realtime',
        'properties.bootstrap.servers' = '',
        'value.format' = 'debezium-json',
        'value.debezium-json.ignore-parse-errors' = 'true',
        'value.debezium-json.timestamp-format.standard' = 'ISO-8601',
        'scan.topic-partition-discovery.interval' = '60s'
    );
{{ end }}

`
	var temp, err = template.New("ddl").Parse(text)
	if err != nil {
		panic(err)
	}
	return &Stg2MysqlHiveDdlTableService{
		configs:     configs,
		endpoint:    endpoint,
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

		row := db.QueryRow(
			"select group_concat(concat('`',"+"COLUMN_NAME,'`', "+`'\t', case
                                                  when DATA_TYPE like '%int' then DATA_TYPE
                                                  when DATA_TYPE like '%char' then COLUMN_TYPE
                                                  when COLUMN_TYPE = 'timestamp' then 'TIMESTAMP_LTZ(3)'
                                                  when DATA_TYPE = 'timestamp'
                                                      then replace(COLUMN_TYPE, 'timestamp', 'TIMESTAMP_LTZ')
                                                  when DATA_TYPE = 'bit' then 'int'
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
