package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HelloHandler struct {
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello\n"))
}

func main() {
	mux := http.NewServeMux()
	terribleSecurityProvider := TerribleSecurityProvider("hello")
	mux.Handle("/hello", terribleSecurityProvider(HelloHandler{}))
	mux.HandleFunc("/world", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("world\n"))
	})
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("greeting!\n"))
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("wang wang\n"))
	})
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))
	server := http.Server{
		Addr:              ":8080",
		Handler:           RequestTimer(mux),
		TLSConfig:         nil,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      90 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		h.ServeHTTP(w, req)
		fmt.Printf("request time %s %v\n", req.URL, time.Since(now))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			handler.ServeHTTP(w, req)
		})
	}
}
