package impl

import (
	"encoding/json"
	"flink-job-exporter/config"
	"flink-job-exporter/dao"
	"flink-job-exporter/service"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Jobs struct {
	Apps struct {
		App []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"app"`
	} `json:"apps"`
}

type FlinkJobServiceImpl struct {
	dao    dao.FlinkJobDao
	config *config.AppConfig
}

func NewFlinkJobServiceImpl(dao dao.FlinkJobDao, config *config.AppConfig) *FlinkJobServiceImpl {
	return &FlinkJobServiceImpl{dao: dao, config: config}
}

func (f *FlinkJobServiceImpl) GetJobStates() []service.JobState {
	prodJobs := f.dao.GetByTag(f.config.Monitor.Tag)
	runningJobs := getRunningJobs(f.config.Yarn.Url)
	var jobs []service.JobState
	for _, job := range prodJobs {
		if has(runningJobs, job) {
			jobs = append(jobs, service.JobState{Name: job, State: 1.0})
		} else {
			jobs = append(jobs, service.JobState{Name: job, State: 0.0})
		}
	}
	return jobs
}

func getRunningJobs(url string) []string {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	body := response.Body
	defer body.Close()
	var jobs Jobs
	bytes, err := io.ReadAll(body)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &jobs)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(jobs)

	var r []string
	for _, app := range jobs.Apps.App {
		r = append(r, app.Name)
	}
	return r
}

func has(jobs []string, job string) bool {
	for _, s := range jobs {
		if s == job {
			return true
		}
	}
	return false
}
