package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func slowServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		select {
		case <-ctx.Done():
			fmt.Println("shutdown slow server")
			return
		case <-time.After(time.Second * 6):
			w.Write([]byte("Slow response"))
		}
	}))
}
func fastServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
}
