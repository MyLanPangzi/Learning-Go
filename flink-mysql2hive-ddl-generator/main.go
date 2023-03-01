package main

import (
	"flag"
	"flink-mysql-ddl-generator/config"
	"flink-mysql-ddl-generator/service"
	"fmt"
	"os"
)

var (
	serviceType string
)

func init() {
	flag.StringVar(&serviceType, "serviceType", "", "mysql2mysql_ddl or ")
}
func main() {
	flag.Parse()

	stgDdlTableService := getService(serviceType)
	err := stgDdlTableService.PrintData(os.Stdout)
	if err != nil {
		panic(err)
	}

}

func getService(serviceType string) service.TableService {
	cfg := config.LoadConfiguration()
	dbService := service.NewDefaultDbService()
	switch serviceType {
	case "mysql2mysql_ddl":
		return service.NewStg2MysqlMysqlDdlTableService(cfg, dbService)
	case "mysql2mysql_dml":
		return service.NewStg2MysqlMysqlDmlTableService(cfg, dbService)
	case "mysql2mysql_flink_hive_ddl":
		return service.NewStg2MysqlHiveDdlTableService(cfg, dbService)
	}
	panic(fmt.Errorf("no such service %v", serviceType))
}
