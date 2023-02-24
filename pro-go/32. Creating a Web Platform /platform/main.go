package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
	"reflect"
)

func main() {
	services.RegisterDefaultServices()
	var cfg config.Configuration
	services.GerService(&cfg)
	var logger logging.Logger
	services.GerService(&logger)
	writeMessage(logger, cfg)
	services.Call(writeMessage)
	s := struct {
		Logger  logging.Logger
		Message string
	}{}
	err := services.PopulateForContextWithExtras(nil, &s, map[reflect.Type]reflect.Value{
		reflect.TypeOf(""): reflect.ValueOf("hello"),
	})
	if err != nil {
		panic(err)
	}
	s.Logger.Info(s.Message)

}
func writeMessage(logger logging.Logger, cfg config.Configuration) {
	message, found := cfg.GetString("main:message")
	if found {
		logger.Info(message)
		return
	}
	logger.Panic("config not found")
}
