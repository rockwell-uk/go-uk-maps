package api

import (
	"net/http"

	"github.com/rockwell-uk/go-logger/logger"

	"go-uk-maps/database"
	"go-uk-maps/httpserver"
)

var (
	HttpPort string = "8080"
	db       database.StorageEngine
)

func Start(dbcn database.StorageEngine) {
	// Assign db to local var
	db = dbcn

	// Whats happening?
	logger.Log(
		logger.LVL_APP,
		"starting api",
	)

	// Apache style logging
	loggingHandler := httpserver.NewLoggingHandler(getMux())

	// Start http
	httpserver.Start(HttpPort, loggingHandler)
}

func Stop() {
	// Whats happening?
	logger.Log(
		logger.LVL_APP,
		"shutting down api",
	)

	// Stop http
	httpserver.Stop()
}

func getMux() *http.ServeMux {
	// Set up http handler
	mux := http.NewServeMux()

	// Favicon
	mux.Handle("/favicon.ico", http.HandlerFunc(faviconHandler))
	mux.Handle("/", http.HandlerFunc(requestHandler))

	return mux
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPost:
		postHandler(w, r)
	}
}
