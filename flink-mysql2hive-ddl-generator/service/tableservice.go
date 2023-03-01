package service

import "io"

type TableService interface {
	PrintData(writer io.Writer) error
}

type TableData struct {
	Db, Table, Columns string
}
