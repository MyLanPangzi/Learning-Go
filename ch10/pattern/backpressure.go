package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}
func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func doThing() string {
	time.Sleep(time.Second * 2)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(writer http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			writer.Write([]byte(doThing()))
		})
		if err != nil {
			writer.WriteHeader(http.StatusTooManyRequests)
			writer.Write([]byte("Too many requests"))
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
