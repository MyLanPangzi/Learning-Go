package main

import (
	"database/sql"
	"flink-job-exporter/config"
	dao "flink-job-exporter/dao/impl"
	"flink-job-exporter/service"
	"flink-job-exporter/service/impl"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func main() {
	c := config.LoadAppConfig()
	fmt.Println(c)
	db := openDb(c)
	defer db.Close()
	var jobDao = dao.NewFlinkJobDaoImpl(db)
	var jobService = impl.NewFlinkJobServiceImpl(jobDao, c)
	registerFlinkJobMetrics(jobService, c)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", c.Http.Host, c.Http.Port), nil))
}

func registerFlinkJobMetrics(jobService service.FlinkJobService, c *config.AppConfig) {
	vec := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "yarn_flink_jobmanager_numRunningJobs",
		Help: "yarn_flink_jobmanager_numRunningJobs",
	}, []string{"job_name"})
	go func() {
		for {
			for _, state := range jobService.GetJobStates() {
				vec.WithLabelValues(state.Name).Set(state.State)
			}
			d, err := time.ParseDuration(c.ScrapeInterval)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(d)
		}
	}()
}

func openDb(c *config.AppConfig) *sql.DB {
	var cfg = mysql.Config{
		User:   c.Db.User,
		Passwd: c.Db.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", c.Db.Host, c.Db.Port),
		DBName: c.Db.Name,
	}
	var db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
