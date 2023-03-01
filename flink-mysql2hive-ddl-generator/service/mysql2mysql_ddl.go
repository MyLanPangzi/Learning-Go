package service

import (
	"flink-mysql-ddl-generator/config"
	"fmt"
	"io"
	"strings"
)

type Stg2MysqlMysqlDdlTableService struct {
	configs   []config.ExportConfig
	endpoint  string
	dbService DbService
}

func NewStg2MysqlMysqlDdlTableService(cfg config.Configuration, dbService DbService) *Stg2MysqlMysqlDdlTableService {
	return &Stg2MysqlMysqlDdlTableService{
		configs:   cfg.Configs,
		endpoint:  cfg.Endpoint,
		dbService: dbService,
	}
}

type Col struct {
	Name, Type, Key string
}

func (c Col) String() string {
	return fmt.Sprintf("%v %v", c.Name, c.Type)
}

var (
	metaCols = []Col{
		{"metadata_schema_name", "varchar(255)", "PRI"},
		{"metadata_table_name", "varchar(255)", "PRI"},
		{"metadata_timestamp", "timestamp(3)", ""},
		{"kafka_topic", "varchar(255)", ""},
		{"kafka_partition", "int", ""},
		{"kafka_offset", "bigint", ""},
		{"kafka_timestamp", "timestamp(3)", ""},
		{"kafka_timestamp_type", "varchar(255)", ""},
	}
)

func metaColsToStrings() []string {
	var out []string
	for _, col := range metaCols {
		out = append(out, col.String())
	}
	return out
}

func (s *Stg2MysqlMysqlDdlTableService) PrintData(w io.Writer) error {
	for _, cfg := range s.configs {
		db, err := s.dbService.OpenDb(cfg, cfg.Db, s.endpoint)
		if err != nil {
			return err
		}
		for _, t := range cfg.Tables {
			ddl := ""
			ignore := ""
			err := db.QueryRow(fmt.Sprintf("show create table %v", t.Regexp)).Scan(&ignore, &ddl)
			if err != nil {
				return err
			}
			r := appendMetaCols(ddl)
			r = removeAutoIncrement(r)
			r = replacePk(r)
			r = appendIfNotExists(r)
			r = replaceTableName(r, t.Table, t.Regexp)
			r += ";\n"

			_, err = w.Write([]byte(r))
			if err != nil {
				return err
			}
		}
		db.Close()
	}
	return nil
}

func replaceTableName(r string, table string, regexp string) string {
	index := strings.Index(r, regexp)
	return r[0:index] + table + r[index+len(regexp):]
}

func appendIfNotExists(r string) string {
	var createTable = "CREATE TABLE"
	index := strings.Index(r, createTable)
	return r[0:index+len(createTable)] + " IF NOT EXISTS " + r[index+len(createTable):]
}

func replacePk(r string) string {
	pkPrefix := "PRIMARY KEY ("
	pkIndex := strings.Index(r, pkPrefix)
	rightIndex := strings.Index(r[pkIndex:], ")")
	return r[0:pkIndex+rightIndex] + ",`metadata_table_name`,`metadata_schema_name`" + r[pkIndex+rightIndex:]
}

func removeAutoIncrement(r string) string {
	var autoIncrement = "AUTO_INCREMENT"
	index := strings.Index(r, autoIncrement)
	if index < 0 {
		return r
	}
	r = r[0:index] + r[index+len(autoIncrement):]
	index = strings.Index(r, "AUTO_INCREMENT=")
	defaultIndex := strings.Index(r[index:], "DEFAULT")

	return r[0:index] + r[index+defaultIndex:]
}

func appendMetaCols(ddl string) string {
	pkPrefix := "PRIMARY KEY ("
	pkIndex := strings.Index(ddl, pkPrefix)
	r := ddl[0:pkIndex]
	r += strings.Join(metaColsToStrings(), ",\n") + ",\n"
	r += ddl[pkIndex:]
	return r
}
