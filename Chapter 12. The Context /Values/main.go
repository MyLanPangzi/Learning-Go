package main

import (
	"Values/controller"
	"Values/identity"
	"Values/logic"
	"Values/tracker"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mux := http.NewServeMux()
		mux.Handle("/process", identity.Middleware(tracker.Middleware(http.HandlerFunc(processUserRequest))))
		logger := tracker.Logger{}
		userController := controller.UserController{
			UserLogic: logic.UserLogic{
				RequestDecorator: tracker.Request,
				Logger:           logger,
				Remote:           "http://localhost:8080/process",
			},
		}
		mux.Handle("/hello", identity.Middleware(tracker.Middleware(userController)))
		server := http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		log.Fatal(server.ListenAndServe())
	}()
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello?data=world", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:   "user",
		Value:  "hello",
		Domain: "localhost:8080",
	})
	response, err := client.Do(tracker.Request(req))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	wg.Wait()

}

func processUserRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	guid, ok := tracker.GuidFromContext(ctx)
	if ok {
		w.Write([]byte(fmt.Sprintf("%s %s", guid, user)))
		return
	}
	w.Write([]byte(fmt.Sprintf("hello %s", user)))
}
