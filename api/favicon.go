package api

import (
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/images/favicon.ico")
}
