package handlers

import (
	"net/http"
)

var klines []int

func (l Handlers) AddStatusHandler(w http.ResponseWriter, r *http.Request) {
	l.writeSuccessResponse(w, map[string]string{
		"success": "simon",
	})
}

func (l Handlers) CurrentStatusHandler(w http.ResponseWriter, r *http.Request) {
	l.writeSuccessResponse(w, klines)
}
