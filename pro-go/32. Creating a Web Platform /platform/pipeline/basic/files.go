package basic

import (
	"net/http"
	"platform/config"
	"platform/pipeline"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
	Cfg           config.Configuration
}

func (s *StaticFileComponent) Init() {
	s.urlPrefix = s.Cfg.GetStringDefault("files:urlprefix", "/files/")
	path, found := s.Cfg.GetString("files:path")
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
