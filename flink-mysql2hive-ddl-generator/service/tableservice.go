package service

import (
	"flink-mysql-ddl-generator/config"
	"fmt"
	"io"
)

type TableService interface {
	PrintData(writer io.Writer) error
}

type TableData struct {
	Db, Table, Columns string
}

func GetService(serviceType string, cfg config.Configuration) TableService {
	dbService := NewDefaultDbService()
	switch serviceType {
	case "mysql2mysql_ddl":
		return NewStg2MysqlMysqlDdlTableService(cfg, dbService)
	case "mysql2mysql_dml":
		return NewStg2MysqlMysqlDmlTableService(cfg)
	case "mysql2mysql_flink_hive_ddl":
		return NewStg2MysqlHiveDdlTableService(cfg, dbService)
	case "mysql2hive":
		return NewStg2HiveTableService(cfg)
	}
	panic(fmt.Errorf("no such service %v", serviceType))
}
