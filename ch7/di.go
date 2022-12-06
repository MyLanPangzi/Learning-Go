package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}
type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}
func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{l: l, ds: ds}
}

func (s SimpleLogic) SayHello(userID string) (string, error) {
	s.l.Log("in SayHello for" + userID)
	name, ok := s.ds.UserNameForId(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "hello " + name, nil
}

func (s SimpleLogic) SayGoodbye(userID string) (string, error) {
	s.l.Log("in SayGoodBye for " + userID)
	name, ok := s.ds.UserNameForId(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "good bye" + name, nil
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{l: l, logic: logic}
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("in SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
