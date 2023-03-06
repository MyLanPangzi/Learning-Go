package service

import (
	"flink-mysql-ddl-generator/config"
	"os"
	"testing"
)

func TestGetService(t *testing.T) {
	type args struct {
		serviceType string
		cfg         config.Configuration
	}
	tests := []struct {
		name string
		args args
		want TableService
	}{
		{name: "mysql2hive ddl", args: args{"mysql2hive", config.LoadConfigurationFrom("testdata/mysql2hive.ddl.json")}, want: nil},
		{name: "mysql2hive dml", args: args{"mysql2hive", config.LoadConfigurationFrom("testdata/mysql2hive.dml.json")}, want: nil},
		{name: "mysql2mysql ddl", args: args{"mysql2mysql_ddl", config.LoadConfigurationFrom("testdata/mysql2mysql.ddl.json")}, want: nil},
		{name: "mysql2mysql dml", args: args{"mysql2mysql_dml", config.LoadConfigurationFrom("testdata/mysql2mysql.dml.json")}, want: nil},
		{name: "mysql2mysql_flink_hive_ddl", args: args{"mysql2mysql_flink_hive_ddl", config.LoadConfigurationFrom("testdata/mysql2mysql_flink_hive.ddl.json")}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := GetService(tt.args.serviceType, tt.args.cfg)
			err := service.PrintData(os.Stdout)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
