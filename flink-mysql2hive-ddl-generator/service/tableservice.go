package service

import "io"

type TableService interface {
	GetData() []TableData
	PrintData(writer io.Writer) error
}

type TableData struct {
	Db, Table, Columns string
}
