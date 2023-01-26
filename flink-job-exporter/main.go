package main

import (
	"database/sql"
	dao "flink-job-exporter/dao/impl"
	"flink-job-exporter/service"
	"flink-job-exporter/service/impl"
	"github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func main() {
	db := openDb()
	defer db.Close()
	var jobDao = dao.NewFlinkJobDaoImpl(db)
	var jobService = impl.NewFlinkJobServiceImpl(jobDao)
	registerFlinkJobMetrics(jobService)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerFlinkJobMetrics(jobService service.FlinkJobService) {
	vec := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "yarn_flink_jobmanager_numRunningJobs",
		Help: "yarn_flink_jobmanager_numRunningJobs",
	}, []string{"job_name"})
	go func() {
		for {
			for _, state := range jobService.GetJobStates() {
				vec.WithLabelValues(state.Name).Set(state.State)
			}
			time.Sleep(time.Second * 3)
		}
	}()
}

func openDb() *sql.DB {
	var cfg = mysql.Config{
		User:   "root",
		Passwd: "!QAZ2wsx",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}
	var db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
