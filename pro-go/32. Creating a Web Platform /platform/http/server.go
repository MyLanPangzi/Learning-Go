package http

import (
	"fmt"
	"net/http"
	"platform/config"
	"platform/logging"
	"platform/pipeline"
	"sync"
	"time"
)

type pipelineAdaptor struct {
	pipeline.RequestPipeline
}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	p.ProcessRequest(request, writer)
}
func Serve(
	pl pipeline.RequestPipeline,
	cfg config.Configuration,
	logger logging.Logger,
) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	adaptor := pipelineAdaptor{pl}
	enableHttp := cfg.GetBoolDefault("http:enableHttp", true)
	if enableHttp {
		httpPort := cfg.GetIntDefault("http:port", 5000)
		logger.Debugf("starting http server on port %v", httpPort)
		wg.Add(1)
		go func() {
			//http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
			server := http.Server{
				Addr:              fmt.Sprintf(":%v", httpPort),
				Handler:           adaptor,
				ReadTimeout:       time.Second * 3,
				ReadHeaderTimeout: time.Second * 1,
				WriteTimeout:      time.Second * 10,
				IdleTimeout:       time.Second * 30,
			}
			panic(server.ListenAndServe())
		}()
	}
	enableHttps := cfg.GetBoolDefault("http:enableHttps", false)
	if enableHttps {
		httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
		logger.Debugf("starting http server on port %v", httpsPort)
		certFile, ok := cfg.GetString("http:httpsCert")
		if !ok {
			panic("https cert not found in settings")
		}
		keyFile, ok := cfg.GetString("http:httpsKey")
		if !ok {
			panic("https key not found in settings")
		}

		wg.Add(1)
		go func() {
			server := http.Server{
				Addr:              fmt.Sprintf(":%v", httpsPort),
				Handler:           adaptor,
				ReadTimeout:       time.Second * 3,
				ReadHeaderTimeout: time.Second * 1,
				WriteTimeout:      time.Second * 10,
				IdleTimeout:       time.Second * 30,
			}
			panic(server.ListenAndServeTLS(certFile, keyFile))
		}()
	}
	return &wg
}
