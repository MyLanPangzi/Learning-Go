package service

type FlinkJobService interface {
	GetJobStates() []JobState
}

type JobState struct {
	Name  string
	State float64
}
