package httpserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const formatPattern = "%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %.4f\n"

type LogRecord struct {
	http.ResponseWriter

	ip                    string
	time                  time.Time
	method, uri, protocol string
	status                int
	responseBytes         int64
	referer               string
	userAgent             string
	elapsedTime           time.Duration
}

//nolint:forbidigo
func (r *LogRecord) Log() {
	timeFormatted := r.time.Format("02/Jan/2006 03:04:05 -0700")
	fmt.Printf(formatPattern, r.ip, timeFormatted, r.method,
		r.uri, r.protocol, r.status, r.responseBytes, r.referer, r.userAgent,
		r.elapsedTime.Seconds())
}

func (r *LogRecord) Write(p []byte) (int, error) {
	written, err := r.ResponseWriter.Write(p)
	r.responseBytes += int64(written)

	return written, err
}

type LoggingHandler struct {
	Handler http.Handler
}

func NewLoggingHandler(handler http.Handler) http.Handler {
	return &LoggingHandler{
		Handler: handler,
	}
}

func (h *LoggingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
		clientIP = clientIP[:colon]
	}

	referer := r.Referer()
	if referer == "" {
		referer = "-"
	}

	userAgent := r.UserAgent()
	if userAgent == "" {
		userAgent = "-"
	}

	record := &LogRecord{
		ResponseWriter: rw,
		ip:             clientIP,
		time:           time.Time{},
		method:         r.Method,
		uri:            r.RequestURI,
		protocol:       r.Proto,
		status:         http.StatusOK,
		referer:        referer,
		userAgent:      userAgent,
		elapsedTime:    time.Duration(0),
	}

	startTime := time.Now()
	h.Handler.ServeHTTP(record, r)
	finishTime := time.Now()

	record.time = finishTime.UTC()
	record.elapsedTime = finishTime.Sub(startTime)

	record.Log()
}
