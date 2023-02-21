package main

import (
	"flag"
	"flink-mysql-ddl-generator/config"
	"flink-mysql-ddl-generator/service"
	"fmt"
	"strings"
)

var (
	mode     string
	endpoint string
)

func init() {
	flag.StringVar(&mode, "mode", "ddl", "--mode ddl or --mode dml. default is ddl. print ddl or dml sql")
	flag.StringVar(&endpoint, "endpoint", "", "aws endpoint")
}
func main() {
	flag.Parse()

	//templates, err := template.ParseFiles("dml.go.html", "ddl.go.html")
	//if err != nil {
	//	panic(err)
	//}
	exportConfig := config.LoadExportConfigs()
	var stgDdlTableService service.TableService = service.NewStg2MysqlMysqlDdlTableService(
		exportConfig,
		endpoint,
		service.DbServiceFunc(service.LocalDbServiceFunc),
	)

	s := &strings.Builder{}
	err := stgDdlTableService.PrintData(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.String())

}
