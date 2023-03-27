package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rockwell-uk/go-logger/logger"

	"go-uk-maps/globvar"
)

var (
	ws   *http.Server
	ctrl = make(chan struct{}, 1)
)

func Start(port string, loggingHandler http.Handler) {
	// Whats happening?
	logger.Log(
		logger.LVL_APP,
		fmt.Sprintf("starting http server on port %s", port),
	)

	// Http server config
	ws = &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           loggingHandler,
		ReadTimeout:       240 * time.Second,
		WriteTimeout:      240 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	// Start the http server
	go func() {
		for {
			err := ws.ListenAndServe()
			if err != nil {
				select {
				case <-ctrl:
					return
				default:
					// Shutdown message
					logger.Log(
						logger.LVL_FATAL,
						err.Error(),
					)
					return
				}
			}
		}
	}()

	logger.Log(
		logger.LVL_APP,
		"started server",
	)
}

func Stop() {
	// Signal close
	ctrl <- struct{}{}

	// Stop http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ws.Shutdown(ctx); err != nil {
		// Handle err
		logger.Log(
			logger.LVL_ERROR,
			fmt.Sprint("error shutting down", err.Error()),
		)
	}

	logger.Log(
		logger.LVL_APP,
		"server exited properly",
	)

	// Report global close success
	globvar.Http <- struct{}{}
}
