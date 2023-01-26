package dao

type FlinkJobDao interface {
	GetByTag(tag string) []string
}
