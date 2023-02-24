package basic

import (
	"net/http"
	"platform/config"
	"platform/pipeline"
	"platform/services"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
}

func (s *StaticFileComponent) Init() {
	var cfg config.Configuration
	err := services.GetService(&cfg)
	if err != nil {
		panic(err)
	}
	s.urlPrefix = cfg.GetStringDefault("files:urlprefix", "/files/")
	path, found := cfg.GetString("files:path")
	if !found {
		panic("cannot load file configuration settings")
	}
	s.stdLibHandler = http.StripPrefix(s.urlPrefix, http.FileServer(http.Dir(path)))
}

func (s *StaticFileComponent) ProcessRequest(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	if !strings.EqualFold(context.URL.Path, s.urlPrefix) &&
		strings.HasPrefix(context.URL.Path, s.urlPrefix) {
		s.stdLibHandler.ServeHTTP(context.ResponseWriter, context.Request)
		return
	}
	next(context)
}

//func (sfc *StaticFileComponent) ProcessRequest(ctx *pipeline.ComponentContext,
//    next func(*pipeline.ComponentContext)) {
//    if  !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
//            strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
//        sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
//    } else {
//        next(ctx)
//    }
//}
