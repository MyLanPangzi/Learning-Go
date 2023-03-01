package service

import (
	"database/sql"
	"encoding/json"
	"flink-mysql-ddl-generator/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/go-sql-driver/mysql"
)

type DbService interface {
	OpenDb(c config.ExportConfig, db, endpoint string) (*sql.DB, error)
}

type DefaultDbService struct {
}

func NewDefaultDbService() *DefaultDbService {
	return &DefaultDbService{}
}

func (a *DefaultDbService) OpenDb(c config.ExportConfig, database, endpoint string) (*sql.DB, error) {
	dbConfig := GetDbConfig(c, endpoint)
	cfg := &mysql.Config{
		User:                 dbConfig.Username,
		Passwd:               dbConfig.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		DBName:               database,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db, err
}

type DbConfig struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

func GetDbConfig(c config.ExportConfig, endpoint string) *DbConfig {
	if c.Username != "" {
		return &DbConfig{
			Username: c.Username,
			Password: c.Password,
			Host:     c.Host,
			Port:     c.Port,
		}
	}
	cfg := aws.NewConfig()
	region := "ap-southeast-1"
	cfg.Region = &region
	if endpoint != "" {
		cfg.Endpoint = &endpoint
	}
	manager := secretsmanager.New(session.Must(session.NewSession(cfg)))
	value, err := manager.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: &c.Arn,
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
