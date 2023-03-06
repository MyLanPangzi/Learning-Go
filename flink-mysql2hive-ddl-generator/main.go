package main

import (
	"flag"
	"flink-mysql-ddl-generator/config"
	"flink-mysql-ddl-generator/service"
	"os"
)

var (
	serviceType string
	template    string
)

func init() {
	flag.StringVar(&serviceType, "serviceType", "", "available service:\n mysql2hive\nmysql2mysql_ddl\nmysql2mysql_dml\nmysql2mysql_flink_hive_ddl  ")
	flag.StringVar(&template, "template", "", "template or ")
}
func main() {
	flag.Parse()

	cfg := config.LoadConfiguration()
	cfg.Template = template

	tableService := service.GetService(serviceType, cfg)
	err := tableService.PrintData(os.Stdout)
	if err != nil {
		panic(err)
	}

}
