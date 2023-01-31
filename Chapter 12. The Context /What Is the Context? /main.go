package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:              ":8080",
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      90 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           Middleware(http.HandlerFunc(handler)),
	}
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "hello", "Hello")
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data := r.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result))
}

func logic(ctx context.Context, data string) (string, error) {
	return ctx.Value("hello").(string) + " " + data, nil
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) call(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://test.com?data"+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	response, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", response.StatusCode)
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(bytes))
	return string(bytes), nil
}
