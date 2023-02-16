package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"text/template"
)

type DbConfig struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}
type ExportConfig struct {
	Arn    string `json:"arn"`
	Db     string `json:"db"`
	Tables []struct {
		Regexp string `json:"regexp"`
		Table  string `json:"table"`
	} `json:"tables"`
}

type Data struct {
	Db, Table, Columns string
}

func NewData(db string, table string, columns string) Data {
	return Data{Db: db, Table: table, Columns: columns}
}

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

	templates, err := template.ParseFiles("dml.go.html", "ddl.go.html")
	if err != nil {
		panic(err)
	}
	exportConfig := loadExportConfigs()
	for _, cfg := range exportConfig {
		var data = makeData(cfg)
		s := &strings.Builder{}
		err := templates.ExecuteTemplate(s, mode+".go.html", data)
		if err != nil {
			panic(err)
		}
		fmt.Println(s.String())
	}

}

func loadExportConfigs() []ExportConfig {
	bytes, err := os.ReadFile("application.json")
	if err != nil {
		panic(err)
	}
	var exportConfig []ExportConfig
	err = json.Unmarshal(bytes, &exportConfig)
	if err != nil {
		panic(err)
	}
	return exportConfig
}

func makeData(cfg ExportConfig) []Data {
	var out []Data
	db, err := openDb(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("set group_concat_max_len =1024*1024*1024 ")
	if err != nil {
		panic(err)
	}
	for _, t := range cfg.Tables {
		//row := db.QueryRow("select table_name from tables where TABLE_SCHEMA = ? and table_name regexp ? limit 1", cfg.Db, t.Regexp)
		//var table string
		//err = row.Scan(&table)
		//if err != nil {
		//	panic(err)
		//}

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
		out = append(out, NewData(cfg.Db, t.Table, columns))
	}
	return out
}

func openDb(config ExportConfig) (*sql.DB, error) {
	dbConfig := getDbConfig(config)
	cfg := &mysql.Config{
		User:                 dbConfig.Username,
		Passwd:               dbConfig.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		DBName:               "information_schema",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db, err
}

func getDbConfig(config ExportConfig) *DbConfig {
	//Local := true
	//if Local {
	//	return &DbConfig{
	//		Username:            " ",
	//		Password:            " ",
	//		Engine:              "mysql",
	//		Host:                "",
	//		Port:                3306,
	//		DbClusterIdentifier: "localhost",
	//	}
	//}
	cfg := aws.NewConfig()
	cfg.Region = stringp("ap-southeast-1")
	if endpoint != "" {
		cfg.Endpoint = &endpoint
	}
	manager := secretsmanager.New(session.Must(session.NewSession(cfg)))
	value, err := manager.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: &config.Arn,
	})
	if err != nil {
		panic(err)
	}
	dbConfig := DbConfig{}
	err = json.Unmarshal([]byte(*value.SecretString), &dbConfig)
	if err != nil {
		panic(err)
	}
	return &dbConfig
}

func stringp(s string) *string {
	return &s
}
